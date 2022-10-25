// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package fixture

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

func easyjsonF4fb9006DecodeGithubComDatabetCloudDatabetGoSdkPkgFixture(in *jlexer.Lexer, out *Competitors) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
	} else {
		in.Delim('{')
		*out = make(Competitors)
		for !in.IsDelim('}') {
			key := string(in.String())
			in.WantColon()
			var v1 Competitor
			(v1).UnmarshalEasyJSON(in)
			(*out)[key] = v1
			in.WantComma()
		}
		in.Delim('}')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF4fb9006EncodeGithubComDatabetCloudDatabetGoSdkPkgFixture(out *jwriter.Writer, in Competitors) {
	if in == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v2First := true
		for v2Name, v2Value := range in {
			if v2First {
				v2First = false
			} else {
				out.RawByte(',')
			}
			out.String(string(v2Name))
			out.RawByte(':')
			(v2Value).MarshalEasyJSON(out)
		}
		out.RawByte('}')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Competitors) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF4fb9006EncodeGithubComDatabetCloudDatabetGoSdkPkgFixture(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Competitors) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF4fb9006EncodeGithubComDatabetCloudDatabetGoSdkPkgFixture(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Competitors) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF4fb9006DecodeGithubComDatabetCloudDatabetGoSdkPkgFixture(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Competitors) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF4fb9006DecodeGithubComDatabetCloudDatabetGoSdkPkgFixture(l, v)
}
func easyjsonF4fb9006DecodeGithubComDatabetCloudDatabetGoSdkPkgFixture1(in *jlexer.Lexer, out *Competitor) {
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
		case "id":
			out.ID = string(in.String())
		case "type":
			out.Type = int(in.Int())
		case "home_away":
			out.HomeAway = int(in.Int())
		case "template_position":
			out.TemplatePosition = int(in.Int())
		case "scores":
			(out.Scores).UnmarshalEasyJSON(in)
		case "name":
			out.Name = string(in.String())
		case "master_id":
			out.MasterID = string(in.String())
		case "country_code":
			out.CountryCode = string(in.String())
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
func easyjsonF4fb9006EncodeGithubComDatabetCloudDatabetGoSdkPkgFixture1(out *jwriter.Writer, in Competitor) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.Int(int(in.Type))
	}
	{
		const prefix string = ",\"home_away\":"
		out.RawString(prefix)
		out.Int(int(in.HomeAway))
	}
	{
		const prefix string = ",\"template_position\":"
		out.RawString(prefix)
		out.Int(int(in.TemplatePosition))
	}
	{
		const prefix string = ",\"scores\":"
		out.RawString(prefix)
		(in.Scores).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"master_id\":"
		out.RawString(prefix)
		out.String(string(in.MasterID))
	}
	{
		const prefix string = ",\"country_code\":"
		out.RawString(prefix)
		out.String(string(in.CountryCode))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Competitor) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF4fb9006EncodeGithubComDatabetCloudDatabetGoSdkPkgFixture1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Competitor) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF4fb9006EncodeGithubComDatabetCloudDatabetGoSdkPkgFixture1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Competitor) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF4fb9006DecodeGithubComDatabetCloudDatabetGoSdkPkgFixture1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Competitor) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF4fb9006DecodeGithubComDatabetCloudDatabetGoSdkPkgFixture1(l, v)
}
