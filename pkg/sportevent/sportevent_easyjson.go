// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package sportevent

import (
	json "encoding/json"
	market "github.com/databet-cloud/databet-go-sdk/pkg/market"
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

func easyjson82c1e1aeDecodeGithubComDatabetCloudDatabetGoSdkPkgSportevent(in *jlexer.Lexer, out *SportEvent) {
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
		case "meta":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.Meta = make(map[string]interface{})
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 interface{}
					if m, ok := v1.(easyjson.Unmarshaler); ok {
						m.UnmarshalEasyJSON(in)
					} else if m, ok := v1.(json.Unmarshaler); ok {
						_ = m.UnmarshalJSON(in.Raw())
					} else {
						v1 = in.Interface()
					}
					(out.Meta)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "fixture":
			(out.Fixture).UnmarshalEasyJSON(in)
		case "MarketIter":
			if in.IsNull() {
				in.Skip()
				out.MarketIter = nil
			} else {
				if out.MarketIter == nil {
					out.MarketIter = new(market.Iterator)
				}
				easyjson82c1e1aeDecodeGithubComDatabetCloudDatabetGoSdkPkgMarket(in, out.MarketIter)
			}
		case "markets":
			(out.Markets).UnmarshalEasyJSON(in)
		case "bet_stop":
			out.BetStop = bool(in.Bool())
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
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
func easyjson82c1e1aeEncodeGithubComDatabetCloudDatabetGoSdkPkgSportevent(out *jwriter.Writer, in SportEvent) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"meta\":"
		out.RawString(prefix)
		if in.Meta == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.Meta {
				if v2First {
					v2First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v2Name))
				out.RawByte(':')
				if m, ok := v2Value.(easyjson.Marshaler); ok {
					m.MarshalEasyJSON(out)
				} else if m, ok := v2Value.(json.Marshaler); ok {
					out.Raw(m.MarshalJSON())
				} else {
					out.Raw(json.Marshal(v2Value))
				}
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"fixture\":"
		out.RawString(prefix)
		(in.Fixture).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"MarketIter\":"
		out.RawString(prefix)
		if in.MarketIter == nil {
			out.RawString("null")
		} else {
			easyjson82c1e1aeEncodeGithubComDatabetCloudDatabetGoSdkPkgMarket(out, *in.MarketIter)
		}
	}
	{
		const prefix string = ",\"markets\":"
		out.RawString(prefix)
		(in.Markets).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"bet_stop\":"
		out.RawString(prefix)
		out.Bool(bool(in.BetStop))
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SportEvent) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson82c1e1aeEncodeGithubComDatabetCloudDatabetGoSdkPkgSportevent(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SportEvent) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson82c1e1aeEncodeGithubComDatabetCloudDatabetGoSdkPkgSportevent(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SportEvent) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson82c1e1aeDecodeGithubComDatabetCloudDatabetGoSdkPkgSportevent(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SportEvent) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson82c1e1aeDecodeGithubComDatabetCloudDatabetGoSdkPkgSportevent(l, v)
}
func easyjson82c1e1aeDecodeGithubComDatabetCloudDatabetGoSdkPkgMarket(in *jlexer.Lexer, out *market.Iterator) {
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
func easyjson82c1e1aeEncodeGithubComDatabetCloudDatabetGoSdkPkgMarket(out *jwriter.Writer, in market.Iterator) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}
