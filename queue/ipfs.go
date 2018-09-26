package queue

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/minio/minio-go"
	log "github.com/sirupsen/logrus"

	"github.com/RTradeLtd/Temporal/mini"

	"github.com/RTradeLtd/Temporal/rtfs"
	"github.com/RTradeLtd/config"

	"github.com/RTradeLtd/database/models"
	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"

	ci "github.com/libp2p/go-libp2p-crypto"
	peer "github.com/libp2p/go-libp2p-peer"
)

// ProcessIPFSKeyCreation is used to create IPFS keys
func (qm *QueueManager) ProcessIPFSKeyCreation(msgs <-chan amqp.Delivery, db *gorm.DB, cfg *config.TemporalConfig) error {
	manager, err := rtfs.Initialize("", "")
	if err != nil {
		return err
	}
	// load the keystore manager
	err = manager.CreateKeystoreManager()
	if err != nil {
		return err
	}
	userManager := models.NewUserManager(db)

	qm.Logger.WithFields(log.Fields{
		"service": qm.QueueName,
	}).Info("processing ipfs key creation requests")

	for d := range msgs {
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
		}).Info("new message detected")

		key := IPFSKeyCreation{}
		err = json.Unmarshal(d.Body, &key)
		if err != nil {
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"error":   err.Error(),
			}).Error("failed to unmarshal message")
			d.Ack(false)
			continue
		}
		if key.NetworkName != "public" {
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"user":    key.UserName,
				"error":   errors.New("private network key creation not yet supported"),
			}).Error("private network key creation not yet supported")
			d.Ack(false)
			continue
		}
		var keyTypeInt int
		var bitsInt int
		switch key.Type {
		case "rsa":
			keyTypeInt = ci.RSA
			if key.Size > 4096 {
				qm.Logger.WithFields(log.Fields{
					"service": qm.QueueName,
					"user":    key.UserName,
					"error":   "key size error",
				}).Error("rsa key generation larger than 4096 bits not supported")
				d.Ack(false)
				continue
			}
			bitsInt = key.Size
		case "ed25519":
			keyTypeInt = ci.Ed25519
			bitsInt = 256
		default:
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"user":    key.UserName,
				"error":   "unsupported key type",
			}).Errorf("%s is not a valid key type, only ed25519 and rsa are supported", key.Type)
			d.Ack(false)
			continue
		}
		keyName := fmt.Sprintf("%s-%s", key.UserName, key.Name)
		pk, err := manager.KeystoreManager.CreateAndSaveKey(keyName, keyTypeInt, bitsInt)
		if err != nil {
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"user":    key.UserName,
				"error":   err.Error(),
			}).Error("failed to create and save key")
			d.Ack(false)
			continue
		}

		id, err := peer.IDFromPrivateKey(pk)
		if err != nil {
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"user":    key.UserName,
				"error":   err.Error(),
			}).Error("failed to get id from private key")
			d.Ack(false)
			continue
		}
		if err := userManager.AddIPFSKeyForUser(key.UserName, keyName, id.Pretty()); err != nil {
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"user":    key.UserName,
				"error":   err.Error(),
			}).Error("failed to add ipfs key to database")
			d.Ack(false)
			continue
		}
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"user":    key.UserName,
		}).Info("successfully processed ipfs key creation")
		d.Ack(false)
	}
	return nil
}

// ProccessIPFSPins is used to process IPFS pin requests
func (qm *QueueManager) ProccessIPFSPins(msgs <-chan amqp.Delivery, db *gorm.DB, cfg *config.TemporalConfig) error {
	userManager := models.NewUserManager(db)
	//uploadManager := models.NewUploadManager(db)
	networkManager := models.NewHostedIPFSNetworkManager(db)
	uploadManager := models.NewUploadManager(db)
	qmEmail, err := Initialize(EmailSendQueue, cfg.RabbitMQ.URL, true, false)
	if err != nil {
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"error":   err.Error(),
		}).Error("failed to initialize email queue connection")
		return err
	}
	qmCluster, err := Initialize(IpfsClusterPinQueue, cfg.RabbitMQ.URL, true, false)
	if err != nil {
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"error":   err.Error(),
		}).Error("failed to initialize cluster pin queue connection")
		return err
	}

	qm.Logger.WithFields(log.Fields{
		"service": qm.QueueName,
	}).Info("processing ipfs pins")

	for d := range msgs {
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
		}).Info("new message detected")

		pin := &IPFSPin{}
		err := json.Unmarshal(d.Body, pin)
		if err != nil {
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"error":   err.Error(),
			}).Error("failed to unmarshal message")
			d.Ack(false)
			continue
		}
		apiURL := ""
		if pin.NetworkName != "public" {
			canAccess, err := userManager.CheckIfUserHasAccessToNetwork(pin.UserName, pin.NetworkName)
			if err != nil {
				qm.Logger.WithFields(log.Fields{
					"service": qm.QueueName,
					"user":    pin.UserName,
					"error":   err.Error(),
				}).Error("error looking up private network in database")
				d.Ack(false)
				continue
			}
			if !canAccess {
				usernames := []string{}
				usernames = append(usernames, pin.UserName)
				es := EmailSend{
					Subject:     IpfsPrivateNetworkUnauthorizedSubject,
					Content:     fmt.Sprintf("Unauthorized access to IPFS private network %s", pin.NetworkName),
					ContentType: "",
					UserNames:   usernames,
				}
				err = qmEmail.PublishMessage(es)
				if err != nil {
					qm.Logger.WithFields(log.Fields{
						"service": qm.QueueName,
						"error":   err.Error(),
					}).Error("failed to publish email send to queue")
				}
				qm.Logger.WithFields(log.Fields{
					"service": qm.QueueName,
					"user":    pin.UserName,
				}).Warn("user does not have access to private network")
				d.Ack(false)
				continue
			}
			url, err := networkManager.GetAPIURLByName(pin.NetworkName)
			if err != nil {
				qm.Logger.WithFields(log.Fields{
					"service": qm.QueueName,
					"user":    pin.UserName,
					"error":   err.Error(),
				}).Error("failed to lookup api url by name in database")
				d.Ack(false)
				continue
			}
			apiURL = url
		}
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"user":    pin.UserName,
		}).Info("initializing connection to IPFS")
		ipfsManager, err := rtfs.Initialize("", apiURL)
		if err != nil {
			addresses := []string{}
			addresses = append(addresses, pin.UserName)
			es := EmailSend{
				Subject:     IpfsInitializationFailedSubject,
				Content:     fmt.Sprintf("Connection to IPFS failed due to the following error %s", err),
				ContentType: "",
				UserNames:   addresses,
			}
			errOne := qmEmail.PublishMessage(es)
			if errOne != nil {
				qm.Logger.WithFields(log.Fields{
					"service": qm.QueueName,
					"error":   err.Error(),
				}).Error("failed to publish email send to queue")
			}
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"user":    pin.UserName,
				"error":   err.Error(),
			}).Error("failed to initialize connection to IPFS")
			d.Ack(false)
			continue
		}
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"user":    pin.UserName,
			"network": pin.NetworkName,
		}).Infof("pinning %s to ipfs", pin.CID)
		err = ipfsManager.Pin(pin.CID)
		if err != nil {
			addresses := []string{}
			addresses = append(addresses, pin.UserName)
			es := EmailSend{
				Subject:     IpfsPinFailedSubject,
				Content:     fmt.Sprintf(IpfsPinFailedContent, pin.CID, pin.NetworkName, err),
				ContentType: "",
				UserNames:   addresses,
			}
			errOne := qmEmail.PublishMessage(es)
			if errOne != nil {
				qm.Logger.WithFields(log.Fields{
					"service": qm.QueueName,
					"error":   err.Error(),
				}).Error("failed to publish email send to queue")
			}
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"user":    pin.UserName,
				"network": pin.NetworkName,
				"error":   err.Error(),
			}).Errorf("failed to pin %s to ipfs", pin.CID)
			d.Ack(false)
			continue
		}
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"user":    pin.UserName,
			"network": pin.NetworkName,
		}).Infof("successfully pinned %s to ipfs", pin.CID)
		clusterAddMsg := IPFSClusterPin{
			CID:              pin.CID,
			NetworkName:      pin.NetworkName,
			HoldTimeInMonths: pin.HoldTimeInMonths,
			UserName:         pin.UserName,
		}
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"user":    pin.UserName,
			"network": pin.NetworkName,
		}).Infof("publishing cluster pin request for %s", pin.CID)
		err = qmCluster.PublishMessage(clusterAddMsg)
		if err != nil {
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"user":    pin.UserName,
				"network": pin.NetworkName,
			}).Errorf("failed to publish cluster pin request for %s", pin.CID)
		}
		_, err = uploadManager.FindUploadByHashAndNetwork(pin.CID, pin.NetworkName)
		if err != nil && err != gorm.ErrRecordNotFound {
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"user":    pin.UserName,
				"network": pin.NetworkName,
				"error":   err.Error(),
			}).Error("failed to find model from database")
			d.Ack(false)
			continue
		}
		if err == gorm.ErrRecordNotFound {
			_, err = uploadManager.NewUpload(pin.CID, "pin", pin.NetworkName, pin.UserName, pin.HoldTimeInMonths)
			if err != nil {
				qm.Logger.WithFields(log.Fields{
					"service": qm.QueueName,
					"user":    pin.UserName,
					"network": pin.NetworkName,
					"error":   err.Error(),
				}).Error("failed to create upload in database")
				d.Ack(false)
				continue
			}
		} else {
			// the record already exists so we will update
			_, err = uploadManager.UpdateUpload(pin.HoldTimeInMonths, pin.UserName, pin.CID, pin.NetworkName)
			if err != nil {
				qm.Logger.WithFields(log.Fields{
					"service": qm.QueueName,
					"user":    pin.UserName,
					"network": pin.NetworkName,
					"error":   err.Error(),
				}).Error("failed to update upload in database")
				d.Ack(false)
				continue
			}
		}
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"user":    pin.UserName,
			"network": pin.NetworkName,
		}).Infof("successfully processed pin for %s", pin.CID)
		d.Ack(false)
	}
	return nil
}

// ProcessIPFSPinRemovals is used to listen for and process any IPFS pin removals.
// This queue must be running on each of the IPFS nodes, and we must eventually run checks
// to ensure that pins were actually removed
func (qm *QueueManager) ProcessIPFSPinRemovals(msgs <-chan amqp.Delivery, cfg *config.TemporalConfig, db *gorm.DB) error {
	userManager := models.NewUserManager(db)
	networkManager := models.NewHostedIPFSNetworkManager(db)
	qmEmail, err := Initialize(EmailSendQueue, cfg.RabbitMQ.URL, true, false)
	if err != nil {
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"error":   err.Error(),
		}).Error("failed to initialize email queue connection")
		return err
	}

	qm.Logger.WithFields(log.Fields{
		"service": qm.QueueName,
	}).Info("processing ipfs pin removals")

	for d := range msgs {
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
		}).Info("detected new message")

		rm := IPFSPinRemoval{}
		err := json.Unmarshal(d.Body, &rm)
		if err != nil {
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"error":   err.Error(),
			}).Error("failed to unmarshal message")
			d.Ack(false)
			continue
		}
		apiURL := ""
		if rm.NetworkName != "public" {
			canAccess, err := userManager.CheckIfUserHasAccessToNetwork(rm.UserName, rm.NetworkName)
			if err != nil {
				qm.Logger.WithFields(log.Fields{
					"service": qm.QueueName,
					"user":    rm.UserName,
					"network": rm.NetworkName,
					"error":   err.Error(),
				}).Error("failed to check database for user network access")
				d.Ack(false)
				continue
			}
			if !canAccess {
				addresses := []string{}
				addresses = append(addresses, rm.UserName)
				es := EmailSend{
					Subject:     IpfsPrivateNetworkUnauthorizedSubject,
					Content:     fmt.Sprintf("Unauthorized access to IPFS private network %s", rm.NetworkName),
					ContentType: "",
					UserNames:   addresses,
				}
				err = qmEmail.PublishMessage(es)
				if err != nil {
					qm.Logger.WithFields(log.Fields{
						"service": qm.QueueName,
						"error":   err.Error(),
					}).Error("failed to publish message to email send queue")
				}
				qm.Logger.WithFields(log.Fields{
					"service": qm.QueueName,
					"user":    rm.UserName,
					"network": rm.NetworkName,
				}).Error("unauthorized access to private network")
				d.Ack(false)
				continue
			}
			apiURL, err = networkManager.GetAPIURLByName(rm.NetworkName)
			if err != nil {
				qm.Logger.WithFields(log.Fields{
					"service": qm.QueueName,
					"user":    rm.UserName,
					"network": rm.NetworkName,
					"error":   err.Error(),
				}).Error("failed to look for api url by name")
				d.Ack(false)
				continue
			}
		}
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"user":    rm.UserName,
			"network": rm.NetworkName,
		}).Info("initializing connection to ipfs")
		ipfsManager, err := rtfs.Initialize("", apiURL)
		if err != nil {
			addresses := []string{rm.UserName}
			es := EmailSend{
				Subject:     IpfsInitializationFailedSubject,
				Content:     fmt.Sprintf("Failed to connect to IPFS network %s for reason %s", rm.NetworkName, err),
				ContentType: "",
				UserNames:   addresses,
			}
			errOne := qmEmail.PublishMessage(es)
			if errOne != nil {
				qm.Logger.WithFields(log.Fields{
					"service": qm.QueueName,
					"error":   errOne.Error(),
				}).Error("failed to publish message to email send queue")
			}
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"user":    rm.UserName,
				"network": rm.NetworkName,
				"error":   err.Error(),
			}).Error("failed to initialize connection to ipfs")
			d.Ack(false)
			continue
		}
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"user":    rm.UserName,
			"network": rm.NetworkName,
		}).Infof("unpinning %s from ipfs", rm.ContentHash)
		err = ipfsManager.Shell.Unpin(rm.ContentHash)
		if err != nil {
			addresses := []string{rm.UserName}
			es := EmailSend{
				Subject:     "Pin removal failed",
				Content:     fmt.Sprintf("Pin removal failed for ipfs network %s due to reason %s", rm.NetworkName, err),
				ContentType: "",
				UserNames:   addresses,
			}
			errOne := qmEmail.PublishMessage(es)
			if errOne != nil {
				qm.Logger.WithFields(log.Fields{
					"service": qm.QueueName,
					"error":   errOne.Error(),
				}).Error("failed to publish message to email send queue")
			}
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"user":    rm.UserName,
				"network": rm.NetworkName,
				"error":   err.Error(),
			}).Errorf("failed to unpin %s", rm.ContentHash)
			d.Ack(false)
			continue
		}
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"user":    rm.UserName,
			"network": rm.NetworkName,
		}).Infof("successfully unpinned %s", rm.ContentHash)
		d.Ack(false)
	}
	return nil
}

// ProccessIPFSFiles is used to process messages sent to rabbitmq to upload files to IPFS.
// This function is invoked with the advanced method of file uploads, and is significantly more resilient than
// the simple file upload method.
func (qm *QueueManager) ProccessIPFSFiles(msgs <-chan amqp.Delivery, cfg *config.TemporalConfig, db *gorm.DB) error {
	// construct the endpoint url to access our minio server
	endpoint := fmt.Sprintf("%s:%s", cfg.MINIO.Connection.IP, cfg.MINIO.Connection.Port)
	// grab our credentials for minio
	accessKey := cfg.MINIO.AccessKey
	secretKey := cfg.MINIO.SecretKey
	ipfsManager, err := rtfs.Initialize("", "")
	if err != nil {
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"error":   err.Error(),
		}).Error("failed to initialize connection to ipfs")
		return err
	}
	// setup our connection to minio
	minioManager, err := mini.NewMinioManager(endpoint, accessKey, secretKey, false)
	if err != nil {
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"error":   err.Error(),
		}).Error("failed to initialize connection to minio")
		return err
	}
	qmEmail, err := Initialize(EmailSendQueue, cfg.RabbitMQ.URL, true, false)
	if err != nil {
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"error":   err.Error(),
		}).Error("failed to initialize email send queue connection")
		return err
	}
	qmPin, err := Initialize(IpfsPinQueue, cfg.RabbitMQ.URL, true, false)
	if err != nil {
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"error":   err.Error(),
		}).Error("failed to initialize pin queue connection")
		return err
	}
	userManager := models.NewUserManager(db)
	networkManager := models.NewHostedIPFSNetworkManager(db)
	uploadManager := models.NewUploadManager(db)
	qm.Logger.WithFields(log.Fields{
		"service": qm.QueueName,
	}).Info("processing ipfs files")
	for d := range msgs {
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
		}).Info("new message detected")

		ipfsFile := IPFSFile{}
		// unmarshal the messagee
		err = json.Unmarshal(d.Body, &ipfsFile)
		if err != nil {
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"error":   err.Error(),
			}).Error("failed to unmarshal message")
			d.Ack(false)
			continue
		}
		if ipfsFile.NetworkName != "public" {
			canAccess, err := userManager.CheckIfUserHasAccessToNetwork(ipfsFile.UserName, ipfsFile.NetworkName)
			if err != nil {
				qm.Logger.WithFields(log.Fields{
					"service": qm.QueueName,
					"user":    ipfsFile.UserName,
					"network": ipfsFile.NetworkName,
					"error":   err.Error(),
				}).Error("failed to check database for user network access")
				d.Ack(false)
				continue
			}
			if !canAccess {
				addresses := []string{}
				addresses = append(addresses, ipfsFile.UserName)
				es := EmailSend{
					Subject:     IpfsPrivateNetworkUnauthorizedSubject,
					Content:     fmt.Sprintf("Unauthorized access to IPFS private network %s", ipfsFile.NetworkName),
					ContentType: "",
					UserNames:   addresses,
				}
				err = qmEmail.PublishMessage(es)
				if err != nil {
					qm.Logger.WithFields(log.Fields{
						"service": qm.QueueName,
						"error":   err.Error(),
					}).Error("failed to publish message to email send queue")
				}
				qm.Logger.WithFields(log.Fields{
					"service": qm.QueueName,
					"user":    ipfsFile.UserName,
					"network": ipfsFile.NetworkName,
				}).Error("unauthorized access to private network")
				d.Ack(false)
				continue
			}
			apiURLName, err := networkManager.GetAPIURLByName(ipfsFile.NetworkName)
			if err != nil {
				qm.Logger.WithFields(log.Fields{
					"service": qm.QueueName,
					"user":    ipfsFile.UserName,
					"network": ipfsFile.NetworkName,
					"error":   err.Error(),
				}).Error("failed to look for api url by name")
				d.Ack(false)
				continue
			}
			apiURL := apiURLName
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"user":    ipfsFile.UserName,
				"network": ipfsFile.NetworkName,
			}).Info("initializing connection to private ipfs network")
			ipfsManager, err = rtfs.Initialize("", apiURL)
			if err != nil {
				addresses := []string{}
				addresses = append(addresses, ipfsFile.UserName)
				es := EmailSend{
					Subject:     IpfsInitializationFailedSubject,
					Content:     fmt.Sprintf("Connection to IPFS failed due to the following error %s", err),
					ContentType: "",
					UserNames:   addresses,
				}
				errOne := qmEmail.PublishMessage(es)
				if errOne != nil {
					qm.Logger.WithFields(log.Fields{
						"service": qm.QueueName,
						"error":   errOne.Error(),
					}).Error("failed to publish message to email send queue")
				}
				qm.Logger.WithFields(log.Fields{
					"service": qm.QueueName,
					"user":    ipfsFile.UserName,
					"network": ipfsFile.NetworkName,
					"error":   err.Error(),
				}).Error("failed to initialize connection to private ipfs network")
				d.Ack(false)
				continue
			}
		}

		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"user":    ipfsFile.UserName,
			"network": ipfsFile.NetworkName,
		}).Info("retrieving object from minio")

		obj, err := minioManager.GetObject(ipfsFile.BucketName, ipfsFile.ObjectName, minio.GetObjectOptions{})
		if err != nil {
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"user":    ipfsFile.UserName,
				"network": ipfsFile.NetworkName,
				"error":   err.Error(),
			}).Info("failed to retrieve object from minio")
			d.Ack(false)
			continue
		}
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"user":    ipfsFile.UserName,
			"network": ipfsFile.NetworkName,
		}).Info("successfully retrieved object from minio")

		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"user":    ipfsFile.UserName,
			"network": ipfsFile.NetworkName,
		}).Info("adding file to ipfs")
		resp, err := ipfsManager.Add(obj)
		if err != nil {
			//TODO: decide how to handle email failures
			addresses := []string{}
			addresses = append(addresses, ipfsFile.UserName)
			es := EmailSend{
				Subject:     IpfsFileFailedSubject,
				Content:     fmt.Sprintf(IpfsFileFailedContent, ipfsFile.ObjectName, ipfsFile.NetworkName),
				ContentType: "",
				UserNames:   addresses,
			}
			errOne := qmEmail.PublishMessage(es)
			if errOne != nil {
				qm.Logger.WithFields(log.Fields{
					"service": qm.QueueName,
					"error":   errOne.Error(),
				}).Error("failed to publish message to email send queue")
			}
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"user":    ipfsFile.UserName,
				"network": ipfsFile.NetworkName,
				"error":   err.Error(),
			}).Info("failed to add file to ipfs")
			d.Ack(false)
			continue
		}

		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"user":    ipfsFile.UserName,
			"network": ipfsFile.NetworkName,
		}).Info("file successfully added to IPFS, forwarding pin request")

		holdTimeInt, err := strconv.ParseInt(ipfsFile.HoldTimeInMonths, 10, 64)
		if err != nil {
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"user":    ipfsFile.UserName,
				"network": ipfsFile.NetworkName,
				"error":   err.Error(),
			}).Warn("failed to parse string to int, using default of 1 month")
			holdTimeInt = 1
		}
		pin := IPFSPin{
			CID:              resp,
			NetworkName:      ipfsFile.NetworkName,
			UserName:         ipfsFile.UserName,
			HoldTimeInMonths: holdTimeInt,
		}

		err = qmPin.PublishMessageWithExchange(pin, PinExchange)
		if err != nil {
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"user":    ipfsFile.UserName,
				"network": ipfsFile.NetworkName,
				"error":   err.Error(),
			}).Warn("failed to publish message to pin queue")
		}

		_, err = uploadManager.FindUploadByHashAndNetwork(resp, ipfsFile.NetworkName)
		if err != nil && err != gorm.ErrRecordNotFound {
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"user":    ipfsFile.UserName,
				"network": ipfsFile.NetworkName,
				"error":   err.Error(),
			}).Error("failed to look for upload in database")
			d.Ack(false)
			continue
		}
		if err == gorm.ErrRecordNotFound {
			_, err = uploadManager.NewUpload(resp, "file", ipfsFile.NetworkName, ipfsFile.UserName, holdTimeInt)
			if err != nil {
				qm.Logger.WithFields(log.Fields{
					"service": qm.QueueName,
					"user":    ipfsFile.UserName,
					"network": ipfsFile.NetworkName,
					"error":   err.Error(),
				}).Error("failed to create new upload in database")
				d.Ack(false)
				continue
			}
		} else {
			_, err = uploadManager.UpdateUpload(holdTimeInt, ipfsFile.UserName, resp, ipfsFile.NetworkName)
			if err != nil {
				qm.Logger.WithFields(log.Fields{
					"service": qm.QueueName,
					"user":    ipfsFile.UserName,
					"network": ipfsFile.NetworkName,
					"error":   err.Error(),
				}).Error("failed to update upload in database")
				d.Ack(false)
				continue
			}
		}
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"user":    ipfsFile.UserName,
			"network": ipfsFile.NetworkName,
		}).Info("removing object from minio")
		err = minioManager.RemoveObject(ipfsFile.BucketName, ipfsFile.ObjectName)
		if err != nil {
			qm.Logger.WithFields(log.Fields{
				"service": qm.QueueName,
				"user":    ipfsFile.UserName,
				"network": ipfsFile.NetworkName,
				"error":   err.Error(),
			}).Info("failed to remove object from minio")
			d.Ack(false)
			continue
		}
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"user":    ipfsFile.UserName,
			"network": ipfsFile.NetworkName,
		}).Info("object removed from minio")
		qm.Logger.WithFields(log.Fields{
			"service": qm.QueueName,
			"user":    ipfsFile.UserName,
			"network": ipfsFile.NetworkName,
		}).Info("succesfully added file into ipfs")
		d.Ack(false)
	}
	return nil
}
