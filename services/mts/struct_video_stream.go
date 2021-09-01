package mts

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// VideoStream is a nested struct in mts response
type VideoStream struct {
	NumFrames      string      `json:"NumFrames" xml:"NumFrames"`
	CodecName      string      `json:"CodecName" xml:"CodecName"`
	Timebase       string      `json:"Timebase" xml:"Timebase"`
	Fps            string      `json:"Fps" xml:"Fps"`
	Sar            string      `json:"Sar" xml:"Sar"`
	PixFmt         string      `json:"PixFmt" xml:"PixFmt"`
	CodecTagString string      `json:"CodecTagString" xml:"CodecTagString"`
	Index          string      `json:"Index" xml:"Index"`
	Lang           string      `json:"Lang" xml:"Lang"`
	HasBFrames     string      `json:"HasBFrames" xml:"HasBFrames"`
	Width          string      `json:"Width" xml:"Width"`
	Rotate         string      `json:"Rotate" xml:"Rotate"`
	CodecTimeBase  string      `json:"CodecTimeBase" xml:"CodecTimeBase"`
	ColorPrimaries string      `json:"ColorPrimaries" xml:"ColorPrimaries"`
	CodecLongName  string      `json:"CodecLongName" xml:"CodecLongName"`
	ColorTransfer  string      `json:"ColorTransfer" xml:"ColorTransfer"`
	AvgFPS         string      `json:"AvgFPS" xml:"AvgFPS"`
	Height         string      `json:"Height" xml:"Height"`
	Dar            string      `json:"Dar" xml:"Dar"`
	Profile        string      `json:"Profile" xml:"Profile"`
	Bitrate        string      `json:"Bitrate" xml:"Bitrate"`
	Level          string      `json:"Level" xml:"Level"`
	CodecTag       string      `json:"CodecTag" xml:"CodecTag"`
	ColorRange     string      `json:"ColorRange" xml:"ColorRange"`
	StartTime      string      `json:"StartTime" xml:"StartTime"`
	Duration       string      `json:"Duration" xml:"Duration"`
	NetworkCost    NetworkCost `json:"NetworkCost" xml:"NetworkCost"`
}
