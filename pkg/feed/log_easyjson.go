// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package feed

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonB6915918DecodeGithubComDatabetCloudDatabetGoSdkPkgFeed(in *jlexer.Lexer, out *LogEntry) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "v":
			out.Version = string(in.String())
		case "sport_event_id":
			out.SportEventID = string(in.String())
		case "type":
			out.Type = LogType(in.String())
		case "timestamp":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
		case "changes":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Patches = make(map[string]json.RawMessage)
				} else {
					out.Patches = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 json.RawMessage
					if data := in.Raw(); in.Ok() {
						in.AddError((v1).UnmarshalJSON(data))
					}
					(out.Patches)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "sport_event":
			(out.SportEvent).UnmarshalEasyJSON(in)
		case "match_id":
			out.MatchID = string(in.String())
		case "market_ids":
			if in.IsNull() {
				in.Skip()
				out.MarketIDs = nil
			} else {
				in.Delim('[')
				if out.MarketIDs == nil {
					if !in.IsDelim(']') {
						out.MarketIDs = make([]string, 0, 4)
					} else {
						out.MarketIDs = []string{}
					}
				} else {
					out.MarketIDs = (out.MarketIDs)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.MarketIDs = append(out.MarketIDs, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "dt_start":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateStart).UnmarshalJSON(data))
			}
		case "dt_end":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateEnd).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB6915918EncodeGithubComDatabetCloudDatabetGoSdkPkgFeed(out *jwriter.Writer, in LogEntry) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"v\":"
		out.RawString(prefix[1:])
		out.String(string(in.Version))
	}
	{
		const prefix string = ",\"sport_event_id\":"
		out.RawString(prefix)
		out.String(string(in.SportEventID))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"timestamp\":"
		out.RawString(prefix)
		easyjsonB6915918EncodeGithubComDatabetCloudDatabetGoSdkPkgFeed1(out, in.Timestamp)
	}
	if len(in.Patches) != 0 {
		const prefix string = ",\"changes\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v3First := true
			for v3Name, v3Value := range in.Patches {
				if v3First {
					v3First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v3Name))
				out.RawByte(':')
				out.Raw((v3Value).MarshalJSON())
			}
			out.RawByte('}')
		}
	}
	if true {
		const prefix string = ",\"sport_event\":"
		out.RawString(prefix)
		(in.SportEvent).MarshalEasyJSON(out)
	}
	if in.MatchID != "" {
		const prefix string = ",\"match_id\":"
		out.RawString(prefix)
		out.String(string(in.MatchID))
	}
	if len(in.MarketIDs) != 0 {
		const prefix string = ",\"market_ids\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v4, v5 := range in.MarketIDs {
				if v4 > 0 {
					out.RawByte(',')
				}
				out.String(string(v5))
			}
			out.RawByte(']')
		}
	}
	if true {
		const prefix string = ",\"dt_start\":"
		out.RawString(prefix)
		out.Raw((in.DateStart).MarshalJSON())
	}
	if true {
		const prefix string = ",\"dt_end\":"
		out.RawString(prefix)
		out.Raw((in.DateEnd).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LogEntry) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB6915918EncodeGithubComDatabetCloudDatabetGoSdkPkgFeed(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LogEntry) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB6915918EncodeGithubComDatabetCloudDatabetGoSdkPkgFeed(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LogEntry) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB6915918DecodeGithubComDatabetCloudDatabetGoSdkPkgFeed(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LogEntry) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB6915918DecodeGithubComDatabetCloudDatabetGoSdkPkgFeed(l, v)
}
func easyjsonB6915918DecodeGithubComDatabetCloudDatabetGoSdkPkgFeed1(in *jlexer.Lexer, out *timestamp) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonB6915918EncodeGithubComDatabetCloudDatabetGoSdkPkgFeed1(out *jwriter.Writer, in timestamp) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}
