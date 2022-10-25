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

func easyjsonAe36e421DecodeGithubComDatabetCloudDatabetGoSdkPkgFixture(in *jlexer.Lexer, out *Platforms) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
	} else {
		in.Delim('{')
		*out = make(Platforms)
		for !in.IsDelim('}') {
			key := string(in.String())
			in.WantColon()
			var v1 Platform
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
func easyjsonAe36e421EncodeGithubComDatabetCloudDatabetGoSdkPkgFixture(out *jwriter.Writer, in Platforms) {
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
func (v Platforms) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAe36e421EncodeGithubComDatabetCloudDatabetGoSdkPkgFixture(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Platforms) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAe36e421EncodeGithubComDatabetCloudDatabetGoSdkPkgFixture(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Platforms) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAe36e421DecodeGithubComDatabetCloudDatabetGoSdkPkgFixture(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Platforms) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAe36e421DecodeGithubComDatabetCloudDatabetGoSdkPkgFixture(l, v)
}
func easyjsonAe36e421DecodeGithubComDatabetCloudDatabetGoSdkPkgFixture1(in *jlexer.Lexer, out *Platform) {
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
		case "type":
			out.Type = string(in.String())
		case "allowed_countries":
			if in.IsNull() {
				in.Skip()
				out.AllowedCountries = nil
			} else {
				in.Delim('[')
				if out.AllowedCountries == nil {
					if !in.IsDelim(']') {
						out.AllowedCountries = make([]string, 0, 4)
					} else {
						out.AllowedCountries = []string{}
					}
				} else {
					out.AllowedCountries = (out.AllowedCountries)[:0]
				}
				for !in.IsDelim(']') {
					var v3 string
					v3 = string(in.String())
					out.AllowedCountries = append(out.AllowedCountries, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "enabled":
			out.Enabled = bool(in.Bool())
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
func easyjsonAe36e421EncodeGithubComDatabetCloudDatabetGoSdkPkgFixture1(out *jwriter.Writer, in Platform) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix[1:])
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"allowed_countries\":"
		out.RawString(prefix)
		if in.AllowedCountries == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v4, v5 := range in.AllowedCountries {
				if v4 > 0 {
					out.RawByte(',')
				}
				out.String(string(v5))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"enabled\":"
		out.RawString(prefix)
		out.Bool(bool(in.Enabled))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Platform) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonAe36e421EncodeGithubComDatabetCloudDatabetGoSdkPkgFixture1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Platform) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonAe36e421EncodeGithubComDatabetCloudDatabetGoSdkPkgFixture1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Platform) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonAe36e421DecodeGithubComDatabetCloudDatabetGoSdkPkgFixture1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Platform) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonAe36e421DecodeGithubComDatabetCloudDatabetGoSdkPkgFixture1(l, v)
}
