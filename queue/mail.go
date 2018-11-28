package queue

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/RTradeLtd/Temporal/mail"
	"github.com/RTradeLtd/config"
	"github.com/streadway/amqp"
)

// ProcessMailSends is a function used to process mail send queue messages
func (qm *Manager) ProcessMailSends(ctx context.Context, wg *sync.WaitGroup, msgs <-chan amqp.Delivery, tCfg *config.TemporalConfig) error {
	mm, err := mail.NewManager(tCfg)
	if err != nil {
		qm.LogError(err, "failed to generate mail manager")
		return err
	}
	qm.LogInfo("processing email sends")
	for d := range msgs {
		qm.LogInfo("detected new message")
		es := EmailSend{}
		if err = json.Unmarshal(d.Body, &es); err != nil {
			qm.LogError(err, "failed to unmarshal message")
			d.Ack(false)
			continue
		}
		var emailSent bool
		for k, v := range es.Emails {
			_, err = mm.SendEmail(es.Subject, es.Content, es.ContentType, es.UserNames[k], v)
			if err != nil {
				qm.LogError(err, "failed to send email")
				d.Ack(false)
				emailSent = false
				continue
			}
			emailSent = true
		}
		if !emailSent {
			continue
		}
		qm.LogInfo("successfully sent email(s)")
		d.Ack(false)
	}
	return nil
}
