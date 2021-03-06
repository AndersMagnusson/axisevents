package axisevents

// collector is passed around to collect the event settings.
type collector struct {
	name            string
	enabled         bool
	version         int
	templateToken   string
	topicExpression string
	messageContent  string
	properties      map[string]interface{}
}
