package control

import "sync"

const MAX_SUBSCRIPTION_ATTEMPTS = 100

type DecodedIndicationMessage struct {
	RequestID             int32
	RequestSequenceNumber int32
	FuncID                int32
	ActionID              int32
	IndSN                 int32
	IndType               int32
	IndHeader             []byte
	IndHeaderLength       int32
	IndMessage            []byte
	IndMessageLength      int32
	CallProcessID         []byte
	CallProcessIDLength   int32
}

type CauseItemType struct {
	CauseType int32
	CauseID   int32
}

type ActionAdmittedListType struct {
	ActionID [16]int32
	Count    int
}

type ActionNotAdmittedListType struct {
	ActionID [16]int32
	Cause    [16]CauseItemType
	Count    int
}

type DecodedSubscriptionResponseMessage struct {
	RequestID             int32
	RequestSequenceNumber int32
	FuncID                int32
	ActionAdmittedList    ActionAdmittedListType
	ActionNotAdmittedList ActionNotAdmittedListType
}

type IntPair64 struct {
	DL int64
	UL int64
}

type OctetString struct {
	Buf  []byte
	Size int
}

type Integer OctetString

type PrintableString OctetString

type ActionDefinition OctetString



type BitString struct {
	Buf        []byte
	Size       int
	BitsUnused int
}

type SubsequentAction struct {
	IsValid              int
	SubsequentActionType int64
	TimeToWait           int64
}

type GNBID BitString

type GlobalgNBIDType struct {
	PlmnID    OctetString
	GnbIDType int
	GnbID     interface{}
}

type GlobalKPMnodegNBIDType struct {
	GlobalgNBID GlobalgNBIDType
	GnbCUUPID   *Integer
	GnbDUID     *Integer
}

type ENGNBID BitString

type GlobalKPMnodeengNBIDType struct {
	PlmnID    OctetString
	GnbIDType int
	GnbID     interface{}
}

type NGENBID_Macro BitString

type NGENBID_ShortMacro BitString

type NGENBID_LongMacro BitString

type GlobalKPMnodengeNBIDType struct {
	PlmnID    OctetString
	EnbIDType int
	EnbID     interface{}
}

type ENBID_Macro BitString

type ENBID_Home BitString

type ENBID_ShortMacro BitString

type ENBID_LongMacro BitString

type GlobalKPMnodeeNBIDType struct {
	PlmnID    OctetString
	EnbIDType int
	EnbID     interface{}
}

type NRCGIType struct {
	PlmnID   OctetString
	NRCellID BitString
}

type SliceIDType struct {
	SST OctetString
	SD  *OctetString
}

type GNB_DU_Name PrintableString

type GNB_CU_CP_Name PrintableString

type GNB_CU_UP_Name PrintableString

type IndicationHeaderFormat1 struct {
	GlobalKPMnodeIDType int32
	GlobalKPMnodeID     interface{}
	NRCGI               *NRCGIType
	PlmnID              *OctetString
	SliceID             *SliceIDType
	FiveQI              int64
	Qci                 int64
	UeMessageType       int32
	GnbDUID             *Integer
	GnbNameType         int32
	GnbName             interface{}
	GlobalgNBID         *GlobalgNBIDType
}

type IndicationHeader struct {
	IndHdrType int32
	IndHdr     interface{}
}

type FQIPERSlicesPerPlmnPerCellType struct {
	FiveQI   int64
	PrbUsage IntPair64
}

type SlicePerPlmnPerCellType struct {
	SliceID                         SliceIDType
	FQIPERSlicesPerPlmnPerCells     [64]FQIPERSlicesPerPlmnPerCellType
	FQIPERSlicesPerPlmnPerCellCount int
}

type DUPM5GCContainerType struct {
	SlicePerPlmnPerCells     [1024]SlicePerPlmnPerCellType
	SlicePerPlmnPerCellCount int
}

type DUPMEPCPerQCIReportType struct {
	QCI      int64
	PrbUsage IntPair64
}

type DUPMEPCContainerType struct {
	PerQCIReports     [256]DUPMEPCPerQCIReportType
	PerQCIReportCount int
}

type ServedPlmnPerCellType struct {
	PlmnID  OctetString
	DUPM5GC *DUPM5GCContainerType
	DUPMEPC *DUPMEPCContainerType
}

type CellResourceReportType struct {
	NRCGI                  NRCGIType
	TotalofAvailablePRBs   IntPair64
	ServedPlmnPerCells     [12]ServedPlmnPerCellType
	ServedPlmnPerCellCount int
}

type ODUPFContainerType struct {
	CellResourceReports     [512]CellResourceReportType
	CellResourceReportCount int
}

type CUCPResourceStatusType struct {
	NumberOfActiveUEs int64
}

type OCUCPPFContainerType struct {
	GNBCUCPName        *PrintableString
	CUCPResourceStatus CUCPResourceStatusType
}

type FQIPERSlicesPerPlmnType struct {
	FiveQI      int64
	PDCPBytesDL *Integer
	PDCPBytesUL *Integer
}

type SliceToReportType struct {
	SliceID                  SliceIDType
	FQIPERSlicesPerPlmns     [64]FQIPERSlicesPerPlmnType
	FQIPERSlicesPerPlmnCount int
}

type CUUPPM5GCType struct {
	SliceToReports     [1024]SliceToReportType
	SliceToReportCount int
}

type CUUPPMEPCPerQCIReportType struct {
	QCI         int64
	PDCPBytesDL *Integer
	PDCPBytesUL *Integer
}

type CUUPPMEPCType struct {
	CUUPPMEPCPerQCIReports     [256]CUUPPMEPCPerQCIReportType
	CUUPPMEPCPerQCIReportCount int
}

type CUUPPlmnType struct {
	PlmnID    OctetString
	CUUPPM5GC *CUUPPM5GCType
	CUUPPMEPC *CUUPPMEPCType
}

type CUUPMeasurementContainerType struct {
	CUUPPlmns     [12]CUUPPlmnType
	CUUPPlmnCount int
}

type CUUPPFContainerItemType struct {
	InterfaceType    int64
	OCUUPPMContainer CUUPMeasurementContainerType
}

type OCUUPPFContainerType struct {
	GNBCUUPName              *PrintableString
	CUUPPFContainerItems     [3]CUUPPFContainerItemType
	CUUPPFContainerItemCount int
}

type DUUsageReportUeResourceReportItemType struct {
	CRNTI      Integer
	PRBUsageDL int64
	PRBUsageUL int64
}

type DUUsageReportCellResourceReportItemType struct {
	NRCGI                     NRCGIType
	UeResourceReportItems     [32]DUUsageReportUeResourceReportItemType
	UeResourceReportItemCount int
}

type DUUsageReportType struct {
	CellResourceReportItems     [512]DUUsageReportCellResourceReportItemType
	CellResourceReportItemCount int
}

type CUCPUsageReportUeResourceReportItemType struct {
	CRNTI          Integer
	ServingCellRF  *OctetString
	NeighborCellRF *OctetString
}

type CUCPUsageReportCellResourceReportItemType struct {
	NRCGI                     NRCGIType
	UeResourceReportItems     [32]CUCPUsageReportUeResourceReportItemType
	UeResourceReportItemCount int
}

type CUCPUsageReportType struct {
	CellResourceReportItems     [16384]CUCPUsageReportCellResourceReportItemType
	CellResourceReportItemCount int
}

type CUUPUsageReportUeResourceReportItemType struct {
	CRNTI       Integer
	PDCPBytesDL *Integer
	PDCPBytesUL *Integer
}

type CUUPUsageReportCellResourceReportItemType struct {
	NRCGI                     NRCGIType
	UeResourceReportItems     [32]CUUPUsageReportUeResourceReportItemType
	UeResourceReportItemCount int
}

type CUUPUsageReportType struct {
	CellResourceReportItems     [512]CUUPUsageReportCellResourceReportItemType
	CellResourceReportItemCount int
}

type PFContainerType struct {
	ContainerType int32
	Container     interface{}
}

type RANContainerType struct {
	Timestamp     OctetString
	ContainerType int32
	Container     interface{}
}

type PMContainerType struct {
	PFContainer  *PFContainerType
	RANContainer *RANContainerType
}
type O_Cu_CpCellMeasurement struct {
	NRCellid						string
	NRCGI           				NRCGIType
	NumberOfSupportedRrcConnections	int
	NumberOfRrcConnections			int
}

type IndicationMessageFormat1 struct {
	o_Cu_CpCellMeasurement [18] O_Cu_CpCellMeasurement   //maxofCells = 18
	ContainerCount   int
}

type cRNTI int
type MeasurementResult struct {
	measurementResultServingCell  MeasurementResultServingCell
}

type MeasurementResultServingCell struct {
	pCI  		int
	basedSSB 	BasedSSB
}
type BasedSSB struct {
	rSRP  		int
	rSRQ 		int
	sINR  		int
}


type O_Cu_CpUeMeasurement struct {
	UEId			int
	NRCellid	 	string
	NRCGI           NRCGIType
	MResult  		MeasurementResult
}

type IndicationMessageFormat4 struct {
	o_Cu_CpUeMeasurement [65535]O_Cu_CpUeMeasurement   //maxofUEs = 65535
	ContainerCount int
}

type IndicationMessage struct {
	StyleType  int64
	IndMsgType int32
	IndMsg     interface{}
}

type Timestamp struct {
	TVsec  int64 `json:"tv_sec"`
	TVnsec int64 `json:"tv_nsec"`
}

type CellMetricsEntry struct {
	MeasTimestampPDCPBytes Timestamp `json:"MeasTimestampPDCPBytes"`
	CellID 		       string 	 `json:"CellID"`
	PDCPBytesDL            int64     `json:"PDCPBytesDL"`
	PDCPBytesUL            int64     `json:"PDCPBytesUL"`
	MeasTimestampPRB       Timestamp `json:"MeasTimestampAvailPRB"`
	AvailPRBDL             int64     `json:"AvailPRBDL"`
	AvailPRBUL             int64     `json:"AvailPRBUL"`
	MeasPeriodPDCP	       int64	 `json:"MeasPeriodPDCPBytes"`
	MeasPeriodPRB	       int64	 `json:"MeasPeriodAvailPRB"`
}

type CellRFType struct {
	RSRP int `json:"rsrp"`
	RSRQ int `json:"rsrq"`
	RSSINR int `json:"rssinr"`
}

type NeighborCellRFType struct {
	CellID string     `json:"CID"`
	CellRF CellRFType `json:"CellRF"`
}

type UeMetricsEntry struct {
	UeID                   int64     `json:"UEID"`
	ServingCellID          string    `json:"ServingCellID"`
	MeasTimestampPDCPBytes Timestamp `json:"MeasTimestampUEPDCPBytes"`
	PDCPBytesDL            int64     `json:"UEPDCPBytesDL"`
	PDCPBytesUL            int64     `json:"UEPDCPBytesUL"`
	MeasTimestampPRB       Timestamp `json:"MeasTimestampUEPRBUsage"`
	PRBUsageDL             int64     `json:"UEPRBUsageDL"`
	PRBUsageUL             int64     `json:"UEPRBUsageUL"`
	MeasTimeRF             Timestamp `json:"MeasTimestampRF"`
	MeasPeriodRF	       int64	 `json:"MeasPeriodRF"`
	MeasPeriodPDCP	       int64	 `json:"MeasPeriodUEPDCPBytes"`
	MeasPeriodPRB	       int64	 `json:"MeasPeriodUEPRBUsage"`
	ServingCellRF   CellRFType           `json:"ServingCellRF"`
	NeighborCellsRF []NeighborCellRFType `json:"NeighborCellRF"`
}

//cnbu add

//为了观察UE切换小区后的位置变化，一段时间内（5分钟，一个模拟周期）假设参与测试的UE数量恒定不变。
type UE struct {
	//dryrun的模拟数据
	Load 			int
	Priority 		int
	SrcCellID 		int    //切换的源cell ；释放的源cell ;实际为 cell数组 的下标
	DestCellID 		int   //切换的目的cell             ;实际为 cell数组 的下标

	//以下为来自基站侧的真实数据
	UEId			int					 //Container.UEId               = int(*Container_C.uEId.c_RNTI)
	MResult  		MeasurementResult    //Container.MResult.measurementResultServingCell.basedSSB.rSRP
}

type Cell struct {
	Cellid	    int					   //从 0 开始分配
	Load 		int
	UEs  		map[int] *UE          // UEId <---> *UE
	SelectUE 	map[int] []int     //Select UEs for Release or HandOver : rSRP <---> UEid
	Lock  		sync.Mutex

	//以下为来自基站侧的真实数据
	NumberOfSupportedRrcConnections	int
	NumberOfRrcConnections			int

	NRCellID		string
	NRCGI           NRCGIType            //Container.NRCGI.PlmnID.Buf   = C.GoBytes(unsafe.Pointer(plmnID_C.buf), C.int(plmnID_C.size))
										//Container.NRCGI.NRCellID.Buf = C.GoBytes(unsafe.Pointer(nRCellID.buf), C.int(nRCellID.size))
}
type GNB struct {
	Cells map[int]*Cell
	CellNum int
	Lock  sync.RWMutex
}


type RICcallProcessID OctetString
type RICcontrolHeader OctetString

