// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package market

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

func easyjson4f6d4409DecodeDatabetGoSdkPkgMarket(in *jlexer.Lexer, out *oddJSON) {
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
		case "template":
			out.Template = string(in.String())
		case "is_active":
			out.IsActive = bool(in.Bool())
		case "status":
			out.Status = OddStatus(in.Int())
		case "value":
			out.Value = float64(in.Float64Str())
		case "marge":
			out.Marge = float64(in.Float64Str())
		case "meta":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Meta).UnmarshalJSON(data))
			}
		case "status_reason":
			out.StatusReason = string(in.String())
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
func easyjson4f6d4409EncodeDatabetGoSdkPkgMarket(out *jwriter.Writer, in oddJSON) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"template\":"
		out.RawString(prefix)
		out.String(string(in.Template))
	}
	{
		const prefix string = ",\"is_active\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsActive))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.Int(int(in.Status))
	}
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		out.Float64Str(float64(in.Value))
	}
	{
		const prefix string = ",\"marge\":"
		out.RawString(prefix)
		out.Float64Str(float64(in.Marge))
	}
	{
		const prefix string = ",\"meta\":"
		out.RawString(prefix)
		out.Raw((in.Meta).MarshalJSON())
	}
	{
		const prefix string = ",\"status_reason\":"
		out.RawString(prefix)
		out.String(string(in.StatusReason))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v oddJSON) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4f6d4409EncodeDatabetGoSdkPkgMarket(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v oddJSON) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4f6d4409EncodeDatabetGoSdkPkgMarket(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *oddJSON) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4f6d4409DecodeDatabetGoSdkPkgMarket(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *oddJSON) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4f6d4409DecodeDatabetGoSdkPkgMarket(l, v)
}
func easyjson4f6d4409DecodeDatabetGoSdkPkgMarket1(in *jlexer.Lexer, out *OddCollection) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
	} else {
		in.Delim('{')
		*out = make(OddCollection)
		for !in.IsDelim('}') {
			key := string(in.String())
			in.WantColon()
			var v1 Odd
			easyjson4f6d4409DecodeDatabetGoSdkPkgMarket2(in, &v1)
			(*out)[key] = v1
			in.WantComma()
		}
		in.Delim('}')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson4f6d4409EncodeDatabetGoSdkPkgMarket1(out *jwriter.Writer, in OddCollection) {
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
func (v OddCollection) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4f6d4409EncodeDatabetGoSdkPkgMarket1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OddCollection) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4f6d4409EncodeDatabetGoSdkPkgMarket1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OddCollection) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4f6d4409DecodeDatabetGoSdkPkgMarket1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OddCollection) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4f6d4409DecodeDatabetGoSdkPkgMarket1(l, v)
}
func easyjson4f6d4409DecodeDatabetGoSdkPkgMarket2(in *jlexer.Lexer, out *Odd) {
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
		case "template":
			out.Template = string(in.String())
		case "is_active":
			out.IsActive = bool(in.Bool())
		case "status":
			out.Status = OddStatus(in.Int())
		case "value":
			out.Value = float64(in.Float64Str())
		case "marge":
			out.Marge = float64(in.Float64Str())
		case "status_reason":
			out.StatusReason = string(in.String())
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
func easyjson4f6d4409EncodeDatabetGoSdkPkgMarket2(out *jwriter.Writer, in Odd) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"template\":"
		out.RawString(prefix)
		out.String(string(in.Template))
	}
	{
		const prefix string = ",\"is_active\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsActive))
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.Int(int(in.Status))
	}
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		out.Float64Str(float64(in.Value))
	}
	{
		const prefix string = ",\"marge\":"
		out.RawString(prefix)
		out.Float64Str(float64(in.Marge))
	}
	{
		const prefix string = ",\"status_reason\":"
		out.RawString(prefix)
		out.String(string(in.StatusReason))
	}
	out.RawByte('}')
}
