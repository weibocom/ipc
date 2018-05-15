package ingest

type Config struct {
	FetchLongTextUrl string `json:"fetch_long_text_url,omitempty" valid:"url,required"`
	LongTextAuth     string `json:"long_text_auth,omitempty" valid:"alphanum,required"`
	PostURL          string `json:"post_url,omitempty" valid:"url,required"`
	Host             string `json:"host,omitempty" valid:"host,required"`
	Port             int    `json:"port,omitempty" valid:"int,required"`
	Consumer         string `json:"consumer,omitempty" valid:"alphanum,required"`
	Retries          int    `json:"retries,omitempty" valid:"int, range(0|10)"`
	ChannelBuffer    int    `json:"channel_buffer,omitempty" valid:"int"`
}
