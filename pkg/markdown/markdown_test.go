package markdown

import (
	"testing"
)

func TestMarkdown(t *testing.T) {
	var err error
	var c = MarkdownContent(content)
	var h []byte
	h, err = c.ToHTML()
	if err != nil {
		t.Error(err)
		return
	}
	if len(content) > len(h) {
		t.Error("Invalid markdown content")
		return
	}

	t.Log(string(h))
}

var content = `
# As Trump refuses to lead, America tries to save itself

(CNN)President Donald Trump isn't leading America much as its pandemic worsens. But that's not stopping Walmart -- along with Kroger, 
Kohl's, and city and state leaders and officials -- from making the tough decisions that the President has shirked.

Given Trump's approach, if the country is to exit the building disaster without many more thousands dead, it will fall to governors, 
mayors, college presidents and school principals, teachers and grocery store managers to execute plans balancing public health with 
the need for life to go on.

There were growing indications Wednesday that such centers of authority across the country are no longer waiting for cues from an 
indifferent President whose aggressive opening strategy has been discredited by a tsunami of infections and whose poll numbers are 
crashing as a result.

The latest Quinnipiac University survey shows Trump trailing presumptive Democratic nominee Joe Biden by 15 points, a deficit that 
might help explain the bizarre series of attacks the President leveled at his rival during Tuesday's news conference and the shake-up 
in his campaign leadership on Wednesday led by his son-in-law and adviser, Jared Kushner.

	One of the NFL's most storied franchises, the Green Bay Packers, will play the preseason without fans. Even Trump's frequent protector, 
	Republican Senate Majority Leader Mitch McConnell, broke with the President's magical thinking as he stumped through his increasingly 
	afflicted home state of Kentucky. McConnell said that while "there were some that hoped" the coronavirus will go away, it isn't
`
