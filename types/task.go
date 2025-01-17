// Types to be shared over the module.
package types

// TaskSettings settings for the task.
type TaskSettings struct {
	Users               []int64 `bson:"users"`
	Url                 string  `bson:"url"`
	CurrentTextTemplate string  `bson:"current_text"`
}

func addDefaultHandler(b *tb.Bot) {
	logger.Println("add default handler")
	b.Handle(tb.OnText, func(c tb.Context) error {
		logger.Println("handle text")
		logger.Println(c.Text())
		return c.Send("unknown command: " + c.Text())
	})

}