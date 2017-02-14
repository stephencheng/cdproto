// AUTOGENERATED FILE: easyjson marshaler/unmarshalers.

package systeminfo

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

func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSysteminfo(in *jlexer.Lexer, out *GPUInfo) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "devices":
			if in.IsNull() {
				in.Skip()
				out.Devices = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.Devices = make([]*GPUDevice, 0, 8)
				} else {
					out.Devices = []*GPUDevice{}
				}
				for !in.IsDelim(']') {
					var v1 *GPUDevice
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(GPUDevice)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Devices = append(out.Devices, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "auxAttributes":
			(out.AuxAttributes).UnmarshalEasyJSON(in)
		case "featureStatus":
			(out.FeatureStatus).UnmarshalEasyJSON(in)
		case "driverBugWorkarounds":
			if in.IsNull() {
				in.Skip()
				out.DriverBugWorkarounds = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.DriverBugWorkarounds = make([]string, 0, 4)
				} else {
					out.DriverBugWorkarounds = []string{}
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.DriverBugWorkarounds = append(out.DriverBugWorkarounds, v2)
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSysteminfo(out *jwriter.Writer, in GPUInfo) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Devices) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"devices\":")
		if in.Devices == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Devices {
				if v3 > 0 {
					out.RawByte(',')
				}
				if v4 == nil {
					out.RawString("null")
				} else {
					(*v4).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if (in.AuxAttributes).IsDefined() {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"auxAttributes\":")
		(in.AuxAttributes).MarshalEasyJSON(out)
	}
	if (in.FeatureStatus).IsDefined() {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"featureStatus\":")
		(in.FeatureStatus).MarshalEasyJSON(out)
	}
	if len(in.DriverBugWorkarounds) != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"driverBugWorkarounds\":")
		if in.DriverBugWorkarounds == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.DriverBugWorkarounds {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GPUInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSysteminfo(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GPUInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSysteminfo(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GPUInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSysteminfo(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GPUInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSysteminfo(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSysteminfo1(in *jlexer.Lexer, out *GPUDevice) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "vendorId":
			out.VendorID = float64(in.Float64())
		case "deviceId":
			out.DeviceID = float64(in.Float64())
		case "vendorString":
			out.VendorString = string(in.String())
		case "deviceString":
			out.DeviceString = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSysteminfo1(out *jwriter.Writer, in GPUDevice) {
	out.RawByte('{')
	first := true
	_ = first
	if in.VendorID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"vendorId\":")
		out.Float64(float64(in.VendorID))
	}
	if in.DeviceID != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"deviceId\":")
		out.Float64(float64(in.DeviceID))
	}
	if in.VendorString != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"vendorString\":")
		out.String(string(in.VendorString))
	}
	if in.DeviceString != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"deviceString\":")
		out.String(string(in.DeviceString))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GPUDevice) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSysteminfo1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GPUDevice) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSysteminfo1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GPUDevice) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSysteminfo1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GPUDevice) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSysteminfo1(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSysteminfo2(in *jlexer.Lexer, out *GetInfoReturns) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "gpu":
			if in.IsNull() {
				in.Skip()
				out.Gpu = nil
			} else {
				if out.Gpu == nil {
					out.Gpu = new(GPUInfo)
				}
				(*out.Gpu).UnmarshalEasyJSON(in)
			}
		case "modelName":
			out.ModelName = string(in.String())
		case "modelVersion":
			out.ModelVersion = string(in.String())
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSysteminfo2(out *jwriter.Writer, in GetInfoReturns) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Gpu != nil {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"gpu\":")
		if in.Gpu == nil {
			out.RawString("null")
		} else {
			(*in.Gpu).MarshalEasyJSON(out)
		}
	}
	if in.ModelName != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"modelName\":")
		out.String(string(in.ModelName))
	}
	if in.ModelVersion != "" {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"modelVersion\":")
		out.String(string(in.ModelVersion))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetInfoReturns) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSysteminfo2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetInfoReturns) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSysteminfo2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetInfoReturns) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSysteminfo2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetInfoReturns) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSysteminfo2(l, v)
}
func easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSysteminfo3(in *jlexer.Lexer, out *GetInfoParams) {
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
		key := in.UnsafeString()
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
func easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSysteminfo3(out *jwriter.Writer, in GetInfoParams) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetInfoParams) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSysteminfo3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetInfoParams) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC5a4559bEncodeGithubComKnqChromedpCdpSysteminfo3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetInfoParams) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSysteminfo3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetInfoParams) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC5a4559bDecodeGithubComKnqChromedpCdpSysteminfo3(l, v)
}
