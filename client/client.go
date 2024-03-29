// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/v2/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type ClassifyingRubbishRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ClassifyingRubbishRequest) String() string {
	return tea.Prettify(s)
}

func (s ClassifyingRubbishRequest) GoString() string {
	return s.String()
}

func (s *ClassifyingRubbishRequest) SetImageURL(v string) *ClassifyingRubbishRequest {
	s.ImageURL = &v
	return s
}

type ClassifyingRubbishAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ClassifyingRubbishAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ClassifyingRubbishAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ClassifyingRubbishAdvanceRequest) SetImageURLObject(v io.Reader) *ClassifyingRubbishAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type ClassifyingRubbishResponseBody struct {
	Data      *ClassifyingRubbishResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClassifyingRubbishResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClassifyingRubbishResponseBody) GoString() string {
	return s.String()
}

func (s *ClassifyingRubbishResponseBody) SetData(v *ClassifyingRubbishResponseBodyData) *ClassifyingRubbishResponseBody {
	s.Data = v
	return s
}

func (s *ClassifyingRubbishResponseBody) SetRequestId(v string) *ClassifyingRubbishResponseBody {
	s.RequestId = &v
	return s
}

type ClassifyingRubbishResponseBodyData struct {
	Elements  []*ClassifyingRubbishResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	Sensitive *bool                                         `json:"Sensitive,omitempty" xml:"Sensitive,omitempty"`
}

func (s ClassifyingRubbishResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ClassifyingRubbishResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClassifyingRubbishResponseBodyData) SetElements(v []*ClassifyingRubbishResponseBodyDataElements) *ClassifyingRubbishResponseBodyData {
	s.Elements = v
	return s
}

func (s *ClassifyingRubbishResponseBodyData) SetSensitive(v bool) *ClassifyingRubbishResponseBodyData {
	s.Sensitive = &v
	return s
}

type ClassifyingRubbishResponseBodyDataElements struct {
	Category      *string  `json:"Category,omitempty" xml:"Category,omitempty"`
	CategoryScore *float32 `json:"CategoryScore,omitempty" xml:"CategoryScore,omitempty"`
	Rubbish       *string  `json:"Rubbish,omitempty" xml:"Rubbish,omitempty"`
	RubbishScore  *float32 `json:"RubbishScore,omitempty" xml:"RubbishScore,omitempty"`
}

func (s ClassifyingRubbishResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s ClassifyingRubbishResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *ClassifyingRubbishResponseBodyDataElements) SetCategory(v string) *ClassifyingRubbishResponseBodyDataElements {
	s.Category = &v
	return s
}

func (s *ClassifyingRubbishResponseBodyDataElements) SetCategoryScore(v float32) *ClassifyingRubbishResponseBodyDataElements {
	s.CategoryScore = &v
	return s
}

func (s *ClassifyingRubbishResponseBodyDataElements) SetRubbish(v string) *ClassifyingRubbishResponseBodyDataElements {
	s.Rubbish = &v
	return s
}

func (s *ClassifyingRubbishResponseBodyDataElements) SetRubbishScore(v float32) *ClassifyingRubbishResponseBodyDataElements {
	s.RubbishScore = &v
	return s
}

type ClassifyingRubbishResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ClassifyingRubbishResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ClassifyingRubbishResponse) String() string {
	return tea.Prettify(s)
}

func (s ClassifyingRubbishResponse) GoString() string {
	return s.String()
}

func (s *ClassifyingRubbishResponse) SetHeaders(v map[string]*string) *ClassifyingRubbishResponse {
	s.Headers = v
	return s
}

func (s *ClassifyingRubbishResponse) SetStatusCode(v int32) *ClassifyingRubbishResponse {
	s.StatusCode = &v
	return s
}

func (s *ClassifyingRubbishResponse) SetBody(v *ClassifyingRubbishResponseBody) *ClassifyingRubbishResponse {
	s.Body = v
	return s
}

type DetectFruitsRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectFruitsRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectFruitsRequest) GoString() string {
	return s.String()
}

func (s *DetectFruitsRequest) SetImageURL(v string) *DetectFruitsRequest {
	s.ImageURL = &v
	return s
}

type DetectFruitsAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s DetectFruitsAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectFruitsAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectFruitsAdvanceRequest) SetImageURLObject(v io.Reader) *DetectFruitsAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type DetectFruitsResponseBody struct {
	Data      *DetectFruitsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectFruitsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectFruitsResponseBody) GoString() string {
	return s.String()
}

func (s *DetectFruitsResponseBody) SetData(v *DetectFruitsResponseBodyData) *DetectFruitsResponseBody {
	s.Data = v
	return s
}

func (s *DetectFruitsResponseBody) SetRequestId(v string) *DetectFruitsResponseBody {
	s.RequestId = &v
	return s
}

type DetectFruitsResponseBodyData struct {
	Elements []*DetectFruitsResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectFruitsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectFruitsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectFruitsResponseBodyData) SetElements(v []*DetectFruitsResponseBodyDataElements) *DetectFruitsResponseBodyData {
	s.Elements = v
	return s
}

type DetectFruitsResponseBodyDataElements struct {
	Box   []*float32 `json:"Box,omitempty" xml:"Box,omitempty" type:"Repeated"`
	Name  *string    `json:"Name,omitempty" xml:"Name,omitempty"`
	Score *float32   `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s DetectFruitsResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectFruitsResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectFruitsResponseBodyDataElements) SetBox(v []*float32) *DetectFruitsResponseBodyDataElements {
	s.Box = v
	return s
}

func (s *DetectFruitsResponseBodyDataElements) SetName(v string) *DetectFruitsResponseBodyDataElements {
	s.Name = &v
	return s
}

func (s *DetectFruitsResponseBodyDataElements) SetScore(v float32) *DetectFruitsResponseBodyDataElements {
	s.Score = &v
	return s
}

type DetectFruitsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectFruitsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectFruitsResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectFruitsResponse) GoString() string {
	return s.String()
}

func (s *DetectFruitsResponse) SetHeaders(v map[string]*string) *DetectFruitsResponse {
	s.Headers = v
	return s
}

func (s *DetectFruitsResponse) SetStatusCode(v int32) *DetectFruitsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectFruitsResponse) SetBody(v *DetectFruitsResponseBody) *DetectFruitsResponse {
	s.Body = v
	return s
}

type DetectImageElementsRequest struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DetectImageElementsRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageElementsRequest) GoString() string {
	return s.String()
}

func (s *DetectImageElementsRequest) SetUrl(v string) *DetectImageElementsRequest {
	s.Url = &v
	return s
}

type DetectImageElementsAdvanceRequest struct {
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DetectImageElementsAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectImageElementsAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectImageElementsAdvanceRequest) SetUrlObject(v io.Reader) *DetectImageElementsAdvanceRequest {
	s.UrlObject = v
	return s
}

type DetectImageElementsResponseBody struct {
	Data      *DetectImageElementsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectImageElementsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectImageElementsResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageElementsResponseBody) SetData(v *DetectImageElementsResponseBodyData) *DetectImageElementsResponseBody {
	s.Data = v
	return s
}

func (s *DetectImageElementsResponseBody) SetRequestId(v string) *DetectImageElementsResponseBody {
	s.RequestId = &v
	return s
}

type DetectImageElementsResponseBodyData struct {
	Elements []*DetectImageElementsResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectImageElementsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectImageElementsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectImageElementsResponseBodyData) SetElements(v []*DetectImageElementsResponseBodyDataElements) *DetectImageElementsResponseBodyData {
	s.Elements = v
	return s
}

type DetectImageElementsResponseBodyDataElements struct {
	Height *int32   `json:"Height,omitempty" xml:"Height,omitempty"`
	Score  *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	Type   *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	Width  *int32   `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *int32   `json:"X,omitempty" xml:"X,omitempty"`
	Y      *int32   `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DetectImageElementsResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s DetectImageElementsResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectImageElementsResponseBodyDataElements) SetHeight(v int32) *DetectImageElementsResponseBodyDataElements {
	s.Height = &v
	return s
}

func (s *DetectImageElementsResponseBodyDataElements) SetScore(v float32) *DetectImageElementsResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectImageElementsResponseBodyDataElements) SetType(v string) *DetectImageElementsResponseBodyDataElements {
	s.Type = &v
	return s
}

func (s *DetectImageElementsResponseBodyDataElements) SetWidth(v int32) *DetectImageElementsResponseBodyDataElements {
	s.Width = &v
	return s
}

func (s *DetectImageElementsResponseBodyDataElements) SetX(v int32) *DetectImageElementsResponseBodyDataElements {
	s.X = &v
	return s
}

func (s *DetectImageElementsResponseBodyDataElements) SetY(v int32) *DetectImageElementsResponseBodyDataElements {
	s.Y = &v
	return s
}

type DetectImageElementsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectImageElementsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectImageElementsResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectImageElementsResponse) GoString() string {
	return s.String()
}

func (s *DetectImageElementsResponse) SetHeaders(v map[string]*string) *DetectImageElementsResponse {
	s.Headers = v
	return s
}

func (s *DetectImageElementsResponse) SetStatusCode(v int32) *DetectImageElementsResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectImageElementsResponse) SetBody(v *DetectImageElementsResponseBody) *DetectImageElementsResponse {
	s.Body = v
	return s
}

type EvaluateCertificateQualityRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s EvaluateCertificateQualityRequest) String() string {
	return tea.Prettify(s)
}

func (s EvaluateCertificateQualityRequest) GoString() string {
	return s.String()
}

func (s *EvaluateCertificateQualityRequest) SetImageURL(v string) *EvaluateCertificateQualityRequest {
	s.ImageURL = &v
	return s
}

func (s *EvaluateCertificateQualityRequest) SetType(v string) *EvaluateCertificateQualityRequest {
	s.Type = &v
	return s
}

type EvaluateCertificateQualityAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Type           *string   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s EvaluateCertificateQualityAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s EvaluateCertificateQualityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *EvaluateCertificateQualityAdvanceRequest) SetImageURLObject(v io.Reader) *EvaluateCertificateQualityAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *EvaluateCertificateQualityAdvanceRequest) SetType(v string) *EvaluateCertificateQualityAdvanceRequest {
	s.Type = &v
	return s
}

type EvaluateCertificateQualityResponseBody struct {
	Data      *EvaluateCertificateQualityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EvaluateCertificateQualityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EvaluateCertificateQualityResponseBody) GoString() string {
	return s.String()
}

func (s *EvaluateCertificateQualityResponseBody) SetData(v *EvaluateCertificateQualityResponseBodyData) *EvaluateCertificateQualityResponseBody {
	s.Data = v
	return s
}

func (s *EvaluateCertificateQualityResponseBody) SetRequestId(v string) *EvaluateCertificateQualityResponseBody {
	s.RequestId = &v
	return s
}

type EvaluateCertificateQualityResponseBodyData struct {
	Elements []*EvaluateCertificateQualityResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s EvaluateCertificateQualityResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EvaluateCertificateQualityResponseBodyData) GoString() string {
	return s.String()
}

func (s *EvaluateCertificateQualityResponseBodyData) SetElements(v []*EvaluateCertificateQualityResponseBodyDataElements) *EvaluateCertificateQualityResponseBodyData {
	s.Elements = v
	return s
}

type EvaluateCertificateQualityResponseBodyDataElements struct {
	Pass  *string `json:"Pass,omitempty" xml:"Pass,omitempty"`
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s EvaluateCertificateQualityResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s EvaluateCertificateQualityResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *EvaluateCertificateQualityResponseBodyDataElements) SetPass(v string) *EvaluateCertificateQualityResponseBodyDataElements {
	s.Pass = &v
	return s
}

func (s *EvaluateCertificateQualityResponseBodyDataElements) SetScore(v string) *EvaluateCertificateQualityResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *EvaluateCertificateQualityResponseBodyDataElements) SetValue(v string) *EvaluateCertificateQualityResponseBodyDataElements {
	s.Value = &v
	return s
}

type EvaluateCertificateQualityResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EvaluateCertificateQualityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EvaluateCertificateQualityResponse) String() string {
	return tea.Prettify(s)
}

func (s EvaluateCertificateQualityResponse) GoString() string {
	return s.String()
}

func (s *EvaluateCertificateQualityResponse) SetHeaders(v map[string]*string) *EvaluateCertificateQualityResponse {
	s.Headers = v
	return s
}

func (s *EvaluateCertificateQualityResponse) SetStatusCode(v int32) *EvaluateCertificateQualityResponse {
	s.StatusCode = &v
	return s
}

func (s *EvaluateCertificateQualityResponse) SetBody(v *EvaluateCertificateQualityResponseBody) *EvaluateCertificateQualityResponse {
	s.Body = v
	return s
}

type GetAsyncJobResultRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAsyncJobResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultRequest) SetJobId(v string) *GetAsyncJobResultRequest {
	s.JobId = &v
	return s
}

type GetAsyncJobResultResponseBody struct {
	Data      *GetAsyncJobResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAsyncJobResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBody) SetData(v *GetAsyncJobResultResponseBodyData) *GetAsyncJobResultResponseBody {
	s.Data = v
	return s
}

func (s *GetAsyncJobResultResponseBody) SetRequestId(v string) *GetAsyncJobResultResponseBody {
	s.RequestId = &v
	return s
}

type GetAsyncJobResultResponseBodyData struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAsyncJobResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorCode(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorMessage(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetJobId(v string) *GetAsyncJobResultResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetResult(v string) *GetAsyncJobResultResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetStatus(v string) *GetAsyncJobResultResponseBodyData {
	s.Status = &v
	return s
}

type GetAsyncJobResultResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAsyncJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAsyncJobResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponse) SetHeaders(v map[string]*string) *GetAsyncJobResultResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncJobResultResponse) SetStatusCode(v int32) *GetAsyncJobResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncJobResultResponse) SetBody(v *GetAsyncJobResultResponseBody) *GetAsyncJobResultResponse {
	s.Body = v
	return s
}

type RecognizeFoodRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeFoodRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFoodRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFoodRequest) SetImageURL(v string) *RecognizeFoodRequest {
	s.ImageURL = &v
	return s
}

type RecognizeFoodAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeFoodAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFoodAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeFoodAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeFoodAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RecognizeFoodResponseBody struct {
	Data      *RecognizeFoodResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeFoodResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFoodResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeFoodResponseBody) SetData(v *RecognizeFoodResponseBodyData) *RecognizeFoodResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeFoodResponseBody) SetRequestId(v string) *RecognizeFoodResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeFoodResponseBodyData struct {
	TopFives []*RecognizeFoodResponseBodyDataTopFives `json:"TopFives,omitempty" xml:"TopFives,omitempty" type:"Repeated"`
}

func (s RecognizeFoodResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFoodResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeFoodResponseBodyData) SetTopFives(v []*RecognizeFoodResponseBodyDataTopFives) *RecognizeFoodResponseBodyData {
	s.TopFives = v
	return s
}

type RecognizeFoodResponseBodyDataTopFives struct {
	Calorie  *string  `json:"Calorie,omitempty" xml:"Calorie,omitempty"`
	Category *string  `json:"Category,omitempty" xml:"Category,omitempty"`
	Score    *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s RecognizeFoodResponseBodyDataTopFives) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFoodResponseBodyDataTopFives) GoString() string {
	return s.String()
}

func (s *RecognizeFoodResponseBodyDataTopFives) SetCalorie(v string) *RecognizeFoodResponseBodyDataTopFives {
	s.Calorie = &v
	return s
}

func (s *RecognizeFoodResponseBodyDataTopFives) SetCategory(v string) *RecognizeFoodResponseBodyDataTopFives {
	s.Category = &v
	return s
}

func (s *RecognizeFoodResponseBodyDataTopFives) SetScore(v float32) *RecognizeFoodResponseBodyDataTopFives {
	s.Score = &v
	return s
}

type RecognizeFoodResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeFoodResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeFoodResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeFoodResponse) GoString() string {
	return s.String()
}

func (s *RecognizeFoodResponse) SetHeaders(v map[string]*string) *RecognizeFoodResponse {
	s.Headers = v
	return s
}

func (s *RecognizeFoodResponse) SetStatusCode(v int32) *RecognizeFoodResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeFoodResponse) SetBody(v *RecognizeFoodResponseBody) *RecognizeFoodResponse {
	s.Body = v
	return s
}

type RecognizeImageColorRequest struct {
	ColorCount *int32  `json:"ColorCount,omitempty" xml:"ColorCount,omitempty"`
	Url        *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecognizeImageColorRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageColorRequest) GoString() string {
	return s.String()
}

func (s *RecognizeImageColorRequest) SetColorCount(v int32) *RecognizeImageColorRequest {
	s.ColorCount = &v
	return s
}

func (s *RecognizeImageColorRequest) SetUrl(v string) *RecognizeImageColorRequest {
	s.Url = &v
	return s
}

type RecognizeImageColorAdvanceRequest struct {
	ColorCount *int32    `json:"ColorCount,omitempty" xml:"ColorCount,omitempty"`
	UrlObject  io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecognizeImageColorAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageColorAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeImageColorAdvanceRequest) SetColorCount(v int32) *RecognizeImageColorAdvanceRequest {
	s.ColorCount = &v
	return s
}

func (s *RecognizeImageColorAdvanceRequest) SetUrlObject(v io.Reader) *RecognizeImageColorAdvanceRequest {
	s.UrlObject = v
	return s
}

type RecognizeImageColorResponseBody struct {
	Data      *RecognizeImageColorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeImageColorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageColorResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeImageColorResponseBody) SetData(v *RecognizeImageColorResponseBodyData) *RecognizeImageColorResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeImageColorResponseBody) SetRequestId(v string) *RecognizeImageColorResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeImageColorResponseBodyData struct {
	ColorTemplateList []*RecognizeImageColorResponseBodyDataColorTemplateList `json:"ColorTemplateList,omitempty" xml:"ColorTemplateList,omitempty" type:"Repeated"`
}

func (s RecognizeImageColorResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageColorResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeImageColorResponseBodyData) SetColorTemplateList(v []*RecognizeImageColorResponseBodyDataColorTemplateList) *RecognizeImageColorResponseBodyData {
	s.ColorTemplateList = v
	return s
}

type RecognizeImageColorResponseBodyDataColorTemplateList struct {
	Color      *string  `json:"Color,omitempty" xml:"Color,omitempty"`
	Label      *string  `json:"Label,omitempty" xml:"Label,omitempty"`
	Percentage *float32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
}

func (s RecognizeImageColorResponseBodyDataColorTemplateList) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageColorResponseBodyDataColorTemplateList) GoString() string {
	return s.String()
}

func (s *RecognizeImageColorResponseBodyDataColorTemplateList) SetColor(v string) *RecognizeImageColorResponseBodyDataColorTemplateList {
	s.Color = &v
	return s
}

func (s *RecognizeImageColorResponseBodyDataColorTemplateList) SetLabel(v string) *RecognizeImageColorResponseBodyDataColorTemplateList {
	s.Label = &v
	return s
}

func (s *RecognizeImageColorResponseBodyDataColorTemplateList) SetPercentage(v float32) *RecognizeImageColorResponseBodyDataColorTemplateList {
	s.Percentage = &v
	return s
}

type RecognizeImageColorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeImageColorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeImageColorResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageColorResponse) GoString() string {
	return s.String()
}

func (s *RecognizeImageColorResponse) SetHeaders(v map[string]*string) *RecognizeImageColorResponse {
	s.Headers = v
	return s
}

func (s *RecognizeImageColorResponse) SetStatusCode(v int32) *RecognizeImageColorResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeImageColorResponse) SetBody(v *RecognizeImageColorResponseBody) *RecognizeImageColorResponse {
	s.Body = v
	return s
}

type RecognizeImageStyleRequest struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecognizeImageStyleRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageStyleRequest) GoString() string {
	return s.String()
}

func (s *RecognizeImageStyleRequest) SetUrl(v string) *RecognizeImageStyleRequest {
	s.Url = &v
	return s
}

type RecognizeImageStyleAdvanceRequest struct {
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RecognizeImageStyleAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageStyleAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeImageStyleAdvanceRequest) SetUrlObject(v io.Reader) *RecognizeImageStyleAdvanceRequest {
	s.UrlObject = v
	return s
}

type RecognizeImageStyleResponseBody struct {
	Data      *RecognizeImageStyleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeImageStyleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageStyleResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeImageStyleResponseBody) SetData(v *RecognizeImageStyleResponseBodyData) *RecognizeImageStyleResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeImageStyleResponseBody) SetRequestId(v string) *RecognizeImageStyleResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeImageStyleResponseBodyData struct {
	Styles []*string `json:"Styles,omitempty" xml:"Styles,omitempty" type:"Repeated"`
}

func (s RecognizeImageStyleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageStyleResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeImageStyleResponseBodyData) SetStyles(v []*string) *RecognizeImageStyleResponseBodyData {
	s.Styles = v
	return s
}

type RecognizeImageStyleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeImageStyleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeImageStyleResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeImageStyleResponse) GoString() string {
	return s.String()
}

func (s *RecognizeImageStyleResponse) SetHeaders(v map[string]*string) *RecognizeImageStyleResponse {
	s.Headers = v
	return s
}

func (s *RecognizeImageStyleResponse) SetStatusCode(v int32) *RecognizeImageStyleResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeImageStyleResponse) SetBody(v *RecognizeImageStyleResponseBody) *RecognizeImageStyleResponse {
	s.Body = v
	return s
}

type RecognizeSceneRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeSceneRequest) GoString() string {
	return s.String()
}

func (s *RecognizeSceneRequest) SetImageURL(v string) *RecognizeSceneRequest {
	s.ImageURL = &v
	return s
}

type RecognizeSceneAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeSceneAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeSceneAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeSceneAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeSceneAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RecognizeSceneResponseBody struct {
	Data      *RecognizeSceneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeSceneResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeSceneResponseBody) SetData(v *RecognizeSceneResponseBodyData) *RecognizeSceneResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeSceneResponseBody) SetRequestId(v string) *RecognizeSceneResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeSceneResponseBodyData struct {
	Tags []*RecognizeSceneResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s RecognizeSceneResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeSceneResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeSceneResponseBodyData) SetTags(v []*RecognizeSceneResponseBodyDataTags) *RecognizeSceneResponseBodyData {
	s.Tags = v
	return s
}

type RecognizeSceneResponseBodyDataTags struct {
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Value      *string  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RecognizeSceneResponseBodyDataTags) String() string {
	return tea.Prettify(s)
}

func (s RecognizeSceneResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *RecognizeSceneResponseBodyDataTags) SetConfidence(v float32) *RecognizeSceneResponseBodyDataTags {
	s.Confidence = &v
	return s
}

func (s *RecognizeSceneResponseBodyDataTags) SetValue(v string) *RecognizeSceneResponseBodyDataTags {
	s.Value = &v
	return s
}

type RecognizeSceneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeSceneResponse) GoString() string {
	return s.String()
}

func (s *RecognizeSceneResponse) SetHeaders(v map[string]*string) *RecognizeSceneResponse {
	s.Headers = v
	return s
}

func (s *RecognizeSceneResponse) SetStatusCode(v int32) *RecognizeSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeSceneResponse) SetBody(v *RecognizeSceneResponseBody) *RecognizeSceneResponse {
	s.Body = v
	return s
}

type RecognizeVehicleTypeRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeVehicleTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleTypeRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleTypeRequest) SetImageURL(v string) *RecognizeVehicleTypeRequest {
	s.ImageURL = &v
	return s
}

type RecognizeVehicleTypeAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RecognizeVehicleTypeAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleTypeAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleTypeAdvanceRequest) SetImageURLObject(v io.Reader) *RecognizeVehicleTypeAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type RecognizeVehicleTypeResponseBody struct {
	Data      *RecognizeVehicleTypeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVehicleTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleTypeResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleTypeResponseBody) SetData(v *RecognizeVehicleTypeResponseBodyData) *RecognizeVehicleTypeResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeVehicleTypeResponseBody) SetRequestId(v string) *RecognizeVehicleTypeResponseBody {
	s.RequestId = &v
	return s
}

type RecognizeVehicleTypeResponseBodyData struct {
	Elements  []*RecognizeVehicleTypeResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	Threshold *float32                                        `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s RecognizeVehicleTypeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleTypeResponseBodyData) SetElements(v []*RecognizeVehicleTypeResponseBodyDataElements) *RecognizeVehicleTypeResponseBodyData {
	s.Elements = v
	return s
}

func (s *RecognizeVehicleTypeResponseBodyData) SetThreshold(v float32) *RecognizeVehicleTypeResponseBodyData {
	s.Threshold = &v
	return s
}

type RecognizeVehicleTypeResponseBodyDataElements struct {
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s RecognizeVehicleTypeResponseBodyDataElements) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleTypeResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleTypeResponseBodyDataElements) SetName(v string) *RecognizeVehicleTypeResponseBodyDataElements {
	s.Name = &v
	return s
}

func (s *RecognizeVehicleTypeResponseBodyDataElements) SetScore(v float32) *RecognizeVehicleTypeResponseBodyDataElements {
	s.Score = &v
	return s
}

type RecognizeVehicleTypeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecognizeVehicleTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecognizeVehicleTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s RecognizeVehicleTypeResponse) GoString() string {
	return s.String()
}

func (s *RecognizeVehicleTypeResponse) SetHeaders(v map[string]*string) *RecognizeVehicleTypeResponse {
	s.Headers = v
	return s
}

func (s *RecognizeVehicleTypeResponse) SetStatusCode(v int32) *RecognizeVehicleTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *RecognizeVehicleTypeResponse) SetBody(v *RecognizeVehicleTypeResponseBody) *RecognizeVehicleTypeResponse {
	s.Body = v
	return s
}

type TaggingAdImageRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s TaggingAdImageRequest) String() string {
	return tea.Prettify(s)
}

func (s TaggingAdImageRequest) GoString() string {
	return s.String()
}

func (s *TaggingAdImageRequest) SetImageURL(v string) *TaggingAdImageRequest {
	s.ImageURL = &v
	return s
}

type TaggingAdImageAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s TaggingAdImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s TaggingAdImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *TaggingAdImageAdvanceRequest) SetImageURLObject(v io.Reader) *TaggingAdImageAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type TaggingAdImageResponseBody struct {
	Data      *TaggingAdImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TaggingAdImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TaggingAdImageResponseBody) GoString() string {
	return s.String()
}

func (s *TaggingAdImageResponseBody) SetData(v *TaggingAdImageResponseBodyData) *TaggingAdImageResponseBody {
	s.Data = v
	return s
}

func (s *TaggingAdImageResponseBody) SetMessage(v string) *TaggingAdImageResponseBody {
	s.Message = &v
	return s
}

func (s *TaggingAdImageResponseBody) SetRequestId(v string) *TaggingAdImageResponseBody {
	s.RequestId = &v
	return s
}

type TaggingAdImageResponseBodyData struct {
	TagInfo map[string]interface{} `json:"TagInfo,omitempty" xml:"TagInfo,omitempty"`
}

func (s TaggingAdImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TaggingAdImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *TaggingAdImageResponseBodyData) SetTagInfo(v map[string]interface{}) *TaggingAdImageResponseBodyData {
	s.TagInfo = v
	return s
}

type TaggingAdImageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TaggingAdImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TaggingAdImageResponse) String() string {
	return tea.Prettify(s)
}

func (s TaggingAdImageResponse) GoString() string {
	return s.String()
}

func (s *TaggingAdImageResponse) SetHeaders(v map[string]*string) *TaggingAdImageResponse {
	s.Headers = v
	return s
}

func (s *TaggingAdImageResponse) SetStatusCode(v int32) *TaggingAdImageResponse {
	s.StatusCode = &v
	return s
}

func (s *TaggingAdImageResponse) SetBody(v *TaggingAdImageResponseBody) *TaggingAdImageResponse {
	s.Body = v
	return s
}

type TaggingImageRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s TaggingImageRequest) String() string {
	return tea.Prettify(s)
}

func (s TaggingImageRequest) GoString() string {
	return s.String()
}

func (s *TaggingImageRequest) SetImageURL(v string) *TaggingImageRequest {
	s.ImageURL = &v
	return s
}

type TaggingImageAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s TaggingImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s TaggingImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *TaggingImageAdvanceRequest) SetImageURLObject(v io.Reader) *TaggingImageAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type TaggingImageResponseBody struct {
	Data      *TaggingImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TaggingImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TaggingImageResponseBody) GoString() string {
	return s.String()
}

func (s *TaggingImageResponseBody) SetData(v *TaggingImageResponseBodyData) *TaggingImageResponseBody {
	s.Data = v
	return s
}

func (s *TaggingImageResponseBody) SetRequestId(v string) *TaggingImageResponseBody {
	s.RequestId = &v
	return s
}

type TaggingImageResponseBodyData struct {
	Tags []*TaggingImageResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s TaggingImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TaggingImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *TaggingImageResponseBodyData) SetTags(v []*TaggingImageResponseBodyDataTags) *TaggingImageResponseBodyData {
	s.Tags = v
	return s
}

type TaggingImageResponseBodyDataTags struct {
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Value      *string  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TaggingImageResponseBodyDataTags) String() string {
	return tea.Prettify(s)
}

func (s TaggingImageResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *TaggingImageResponseBodyDataTags) SetConfidence(v float32) *TaggingImageResponseBodyDataTags {
	s.Confidence = &v
	return s
}

func (s *TaggingImageResponseBodyDataTags) SetValue(v string) *TaggingImageResponseBodyDataTags {
	s.Value = &v
	return s
}

type TaggingImageResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TaggingImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TaggingImageResponse) String() string {
	return tea.Prettify(s)
}

func (s TaggingImageResponse) GoString() string {
	return s.String()
}

func (s *TaggingImageResponse) SetHeaders(v map[string]*string) *TaggingImageResponse {
	s.Headers = v
	return s
}

func (s *TaggingImageResponse) SetStatusCode(v int32) *TaggingImageResponse {
	s.StatusCode = &v
	return s
}

func (s *TaggingImageResponse) SetBody(v *TaggingImageResponseBody) *TaggingImageResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("imagerecog"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClassifyingRubbishWithOptions(request *ClassifyingRubbishRequest, runtime *util.RuntimeOptions) (_result *ClassifyingRubbishResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ClassifyingRubbish"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ClassifyingRubbishResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClassifyingRubbish(request *ClassifyingRubbishRequest) (_result *ClassifyingRubbishResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClassifyingRubbishResponse{}
	_body, _err := client.ClassifyingRubbishWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClassifyingRubbishAdvance(request *ClassifyingRubbishAdvanceRequest, runtime *util.RuntimeOptions) (_result *ClassifyingRubbishResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imagerecog"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	classifyingRubbishReq := &ClassifyingRubbishRequest{}
	openapiutil.Convert(request, classifyingRubbishReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		classifyingRubbishReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	classifyingRubbishResp, _err := client.ClassifyingRubbishWithOptions(classifyingRubbishReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = classifyingRubbishResp
	return _result, _err
}

func (client *Client) DetectFruitsWithOptions(request *DetectFruitsRequest, runtime *util.RuntimeOptions) (_result *DetectFruitsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectFruits"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectFruitsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectFruits(request *DetectFruitsRequest) (_result *DetectFruitsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectFruitsResponse{}
	_body, _err := client.DetectFruitsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectFruitsAdvance(request *DetectFruitsAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectFruitsResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imagerecog"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectFruitsReq := &DetectFruitsRequest{}
	openapiutil.Convert(request, detectFruitsReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectFruitsReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectFruitsResp, _err := client.DetectFruitsWithOptions(detectFruitsReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectFruitsResp
	return _result, _err
}

func (client *Client) DetectImageElementsWithOptions(request *DetectImageElementsRequest, runtime *util.RuntimeOptions) (_result *DetectImageElementsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectImageElements"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectImageElementsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectImageElements(request *DetectImageElementsRequest) (_result *DetectImageElementsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectImageElementsResponse{}
	_body, _err := client.DetectImageElementsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectImageElementsAdvance(request *DetectImageElementsAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectImageElementsResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imagerecog"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	detectImageElementsReq := &DetectImageElementsRequest{}
	openapiutil.Convert(request, detectImageElementsReq)
	if !tea.BoolValue(util.IsUnset(request.UrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.UrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		detectImageElementsReq.Url = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	detectImageElementsResp, _err := client.DetectImageElementsWithOptions(detectImageElementsReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectImageElementsResp
	return _result, _err
}

func (client *Client) EvaluateCertificateQualityWithOptions(request *EvaluateCertificateQualityRequest, runtime *util.RuntimeOptions) (_result *EvaluateCertificateQualityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EvaluateCertificateQuality"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EvaluateCertificateQualityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EvaluateCertificateQuality(request *EvaluateCertificateQualityRequest) (_result *EvaluateCertificateQualityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EvaluateCertificateQualityResponse{}
	_body, _err := client.EvaluateCertificateQualityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EvaluateCertificateQualityAdvance(request *EvaluateCertificateQualityAdvanceRequest, runtime *util.RuntimeOptions) (_result *EvaluateCertificateQualityResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imagerecog"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	evaluateCertificateQualityReq := &EvaluateCertificateQualityRequest{}
	openapiutil.Convert(request, evaluateCertificateQualityReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		evaluateCertificateQualityReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	evaluateCertificateQualityResp, _err := client.EvaluateCertificateQualityWithOptions(evaluateCertificateQualityReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = evaluateCertificateQualityResp
	return _result, _err
}

func (client *Client) GetAsyncJobResultWithOptions(request *GetAsyncJobResultRequest, runtime *util.RuntimeOptions) (_result *GetAsyncJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsyncJobResult"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncJobResult(request *GetAsyncJobResultRequest) (_result *GetAsyncJobResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.GetAsyncJobResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeFoodWithOptions(request *RecognizeFoodRequest, runtime *util.RuntimeOptions) (_result *RecognizeFoodResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeFood"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeFoodResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeFood(request *RecognizeFoodRequest) (_result *RecognizeFoodResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeFoodResponse{}
	_body, _err := client.RecognizeFoodWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeFoodAdvance(request *RecognizeFoodAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeFoodResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imagerecog"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeFoodReq := &RecognizeFoodRequest{}
	openapiutil.Convert(request, recognizeFoodReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeFoodReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeFoodResp, _err := client.RecognizeFoodWithOptions(recognizeFoodReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeFoodResp
	return _result, _err
}

func (client *Client) RecognizeImageColorWithOptions(request *RecognizeImageColorRequest, runtime *util.RuntimeOptions) (_result *RecognizeImageColorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ColorCount)) {
		body["ColorCount"] = request.ColorCount
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeImageColor"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeImageColorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeImageColor(request *RecognizeImageColorRequest) (_result *RecognizeImageColorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeImageColorResponse{}
	_body, _err := client.RecognizeImageColorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeImageColorAdvance(request *RecognizeImageColorAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeImageColorResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imagerecog"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeImageColorReq := &RecognizeImageColorRequest{}
	openapiutil.Convert(request, recognizeImageColorReq)
	if !tea.BoolValue(util.IsUnset(request.UrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.UrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeImageColorReq.Url = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeImageColorResp, _err := client.RecognizeImageColorWithOptions(recognizeImageColorReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeImageColorResp
	return _result, _err
}

func (client *Client) RecognizeImageStyleWithOptions(request *RecognizeImageStyleRequest, runtime *util.RuntimeOptions) (_result *RecognizeImageStyleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeImageStyle"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeImageStyleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeImageStyle(request *RecognizeImageStyleRequest) (_result *RecognizeImageStyleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeImageStyleResponse{}
	_body, _err := client.RecognizeImageStyleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeImageStyleAdvance(request *RecognizeImageStyleAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeImageStyleResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imagerecog"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeImageStyleReq := &RecognizeImageStyleRequest{}
	openapiutil.Convert(request, recognizeImageStyleReq)
	if !tea.BoolValue(util.IsUnset(request.UrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.UrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeImageStyleReq.Url = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeImageStyleResp, _err := client.RecognizeImageStyleWithOptions(recognizeImageStyleReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeImageStyleResp
	return _result, _err
}

func (client *Client) RecognizeSceneWithOptions(request *RecognizeSceneRequest, runtime *util.RuntimeOptions) (_result *RecognizeSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeScene"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeScene(request *RecognizeSceneRequest) (_result *RecognizeSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeSceneResponse{}
	_body, _err := client.RecognizeSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeSceneAdvance(request *RecognizeSceneAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeSceneResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imagerecog"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeSceneReq := &RecognizeSceneRequest{}
	openapiutil.Convert(request, recognizeSceneReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeSceneReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeSceneResp, _err := client.RecognizeSceneWithOptions(recognizeSceneReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeSceneResp
	return _result, _err
}

func (client *Client) RecognizeVehicleTypeWithOptions(request *RecognizeVehicleTypeRequest, runtime *util.RuntimeOptions) (_result *RecognizeVehicleTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RecognizeVehicleType"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecognizeVehicleTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecognizeVehicleType(request *RecognizeVehicleTypeRequest) (_result *RecognizeVehicleTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecognizeVehicleTypeResponse{}
	_body, _err := client.RecognizeVehicleTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecognizeVehicleTypeAdvance(request *RecognizeVehicleTypeAdvanceRequest, runtime *util.RuntimeOptions) (_result *RecognizeVehicleTypeResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imagerecog"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	recognizeVehicleTypeReq := &RecognizeVehicleTypeRequest{}
	openapiutil.Convert(request, recognizeVehicleTypeReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		recognizeVehicleTypeReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	recognizeVehicleTypeResp, _err := client.RecognizeVehicleTypeWithOptions(recognizeVehicleTypeReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = recognizeVehicleTypeResp
	return _result, _err
}

func (client *Client) TaggingAdImageWithOptions(request *TaggingAdImageRequest, runtime *util.RuntimeOptions) (_result *TaggingAdImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TaggingAdImage"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TaggingAdImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TaggingAdImage(request *TaggingAdImageRequest) (_result *TaggingAdImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TaggingAdImageResponse{}
	_body, _err := client.TaggingAdImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TaggingAdImageAdvance(request *TaggingAdImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *TaggingAdImageResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imagerecog"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	taggingAdImageReq := &TaggingAdImageRequest{}
	openapiutil.Convert(request, taggingAdImageReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		taggingAdImageReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	taggingAdImageResp, _err := client.TaggingAdImageWithOptions(taggingAdImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = taggingAdImageResp
	return _result, _err
}

func (client *Client) TaggingImageWithOptions(request *TaggingImageRequest, runtime *util.RuntimeOptions) (_result *TaggingImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TaggingImage"),
		Version:     tea.String("2019-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TaggingImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TaggingImage(request *TaggingImageRequest) (_result *TaggingImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TaggingImageResponse{}
	_body, _err := client.TaggingImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TaggingImageAdvance(request *TaggingImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *TaggingImageResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("imagerecog"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	taggingImageReq := &TaggingImageRequest{}
	openapiutil.Convert(request, taggingImageReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		taggingImageReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	taggingImageResp, _err := client.TaggingImageWithOptions(taggingImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = taggingImageResp
	return _result, _err
}
