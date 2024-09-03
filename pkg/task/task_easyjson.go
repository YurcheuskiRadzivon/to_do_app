// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package task

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

func easyjson79a0a577DecodeGithubComYurcheuskiRadzivonToDoAppPkgTask(in *jlexer.Lexer, out *Tasks) {
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
		case "list":
			if in.IsNull() {
				in.Skip()
				out.List = nil
			} else {
				in.Delim('[')
				if out.List == nil {
					if !in.IsDelim(']') {
						out.List = make([]Task, 0, 1)
					} else {
						out.List = []Task{}
					}
				} else {
					out.List = (out.List)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Task
					(v1).UnmarshalEasyJSON(in)
					out.List = append(out.List, v1)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson79a0a577EncodeGithubComYurcheuskiRadzivonToDoAppPkgTask(out *jwriter.Writer, in Tasks) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"list\":"
		out.RawString(prefix[1:])
		if in.List == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.List {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Tasks) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson79a0a577EncodeGithubComYurcheuskiRadzivonToDoAppPkgTask(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Tasks) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson79a0a577DecodeGithubComYurcheuskiRadzivonToDoAppPkgTask(l, v)
}
func easyjson79a0a577DecodeGithubComYurcheuskiRadzivonToDoAppPkgTask1(in *jlexer.Lexer, out *Task) {
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
			out.ID = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "notes":
			out.Notes = string(in.String())
		case "completed":
			out.Completed = bool(in.Bool())
		case "priority":
			out.Priority = int(in.Int())
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
func easyjson79a0a577EncodeGithubComYurcheuskiRadzivonToDoAppPkgTask1(out *jwriter.Writer, in Task) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"notes\":"
		out.RawString(prefix)
		out.String(string(in.Notes))
	}
	{
		const prefix string = ",\"completed\":"
		out.RawString(prefix)
		out.Bool(bool(in.Completed))
	}
	{
		const prefix string = ",\"priority\":"
		out.RawString(prefix)
		out.Int(int(in.Priority))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Task) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson79a0a577EncodeGithubComYurcheuskiRadzivonToDoAppPkgTask1(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Task) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson79a0a577DecodeGithubComYurcheuskiRadzivonToDoAppPkgTask1(l, v)
}
