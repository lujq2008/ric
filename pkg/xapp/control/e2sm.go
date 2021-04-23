
package control

// #cgo CFLAGS: -I../../../internal/asn1codec/ -DASN_DISABLE_OER_SUPPORT
// #cgo LDFLAGS: -L../../../internal/asn1codec/   -lasn1objects
/*
#include <wrapper_e2sm.h>
*/
import "C"

import (
	"bytes"
	"encoding/binary"
	"errors"
	"strconv"
	"unsafe"
)

type E2sm struct {
}

func (c *E2sm) SetEventTriggerDefinition(buffer []byte, eventTriggerCount int, RTPeriods []int64) (newBuffer []byte, err error) {
	cptr := unsafe.Pointer(&buffer[0])
	periods := unsafe.Pointer(&RTPeriods[0])
	size := C.e2sm_encode_ric_event_trigger_definition(cptr, C.size_t(len(buffer)), C.size_t(eventTriggerCount), (*C.long)(periods))
	if size < 0 {
		return make([]byte, 0), errors.New("e2sm wrapper is unable to set EventTriggerDefinition due to wrong or invalid input")
	}
	newBuffer = C.GoBytes(cptr, (C.int(size)+7)/8)
	return
}

func (c *E2sm) SetActionDefinition(buffer []byte, ricStyleType int64) (newBuffer []byte, err error) {
	cptr := unsafe.Pointer(&buffer[0])
	size := C.e2sm_encode_ric_action_definition(cptr, C.size_t(len(buffer)), C.long(ricStyleType))
	if size < 0 {
		return make([]byte, 0), errors.New("e2sm wrapper is unable to set ActionDefinition due to wrong or invalid input")
	}
	newBuffer = C.GoBytes(cptr, (C.int(size)+7)/8)
	return
}


func (c *E2sm) SetControlHeader(buffer []byte, ricStyleType int64) (newBuffer []byte, err error) {
	cptr := unsafe.Pointer(&buffer[0])
	size := C.e2sm_encode_ric_control_header(cptr, C.size_t(len(buffer)), C.long(ricStyleType))
	if size < 0 {
		return make([]byte, 0), errors.New("e2sm wrapper is unable to set ControlHeader due to wrong or invalid input")
	}
	newBuffer = C.GoBytes(cptr, (C.int(size)+7)/8)
	return
}

//var mutex_SetControlMessageFormat6                sync.Mutex
func (c *E2sm) SetControlMessageFormat6(buffer []byte, ReleaseUEs map[int]*UE, HandOverUEs map[int]*UE) (newBuffer []byte, err error) {
	//mutex_SetControlMessageFormat6.Lock()
	//defer mutex_SetControlMessageFormat6.Unlock()

	ReleaseCommandPtrs := (**C.ReleaseCommand_t)(C.calloc(C.size_t(len(ReleaseUEs)), 8))
	defer C.free(unsafe.Pointer(ReleaseCommandPtrs))
	index := 0
	for _,v := range ReleaseUEs {
		var c_RNTI C.long
		var pLMNIdentitySrc C.OCTET_STRING_t
		var nRCellIdentitySrc C.BIT_STRING_t

		c_RNTI = C.long(v.UEId)
		pLMNIdentitySrc.buf = (*C.uint8_t)(C.CBytes(GnoteB.Cells[v.SrcCellID].NRCGI.PlmnID.Buf))
		pLMNIdentitySrc.size = C.ulong(GnoteB.Cells[v.SrcCellID].NRCGI.PlmnID.Size)
		//defer C.free(unsafe.Pointer(pLMNIdentitySrc.buf))
		//fmt.Printf("pLMNIdentitySrc.buf = %v\n",pLMNIdentitySrc.buf)

		nRCellIdentitySrc.buf 			= (*C.uint8_t)(C.CBytes(GnoteB.Cells[v.SrcCellID].NRCGI.NRCellID.Buf))
		nRCellIdentitySrc.size 			= C.ulong(GnoteB.Cells[v.SrcCellID].NRCGI.NRCellID.Size)
		nRCellIdentitySrc.bits_unused 	= C.int(GnoteB.Cells[v.SrcCellID].NRCGI.NRCellID.BitsUnused)
		//defer C.free(unsafe.Pointer(nRCellIdentitySrc.buf))


		ReleaseCommand := C.e2sm_encode_ric_control_release_command(pLMNIdentitySrc, nRCellIdentitySrc, c_RNTI)
		*(**C.ReleaseCommand_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(ReleaseCommandPtrs)) + (uintptr)(8*C.int(index)))) = ReleaseCommand

		//ReleaseCommand_slice = append(ReleaseCommand_slice,ReleaseCommand)
		index++
		/*
		fmt.Printf("pLMNIdentitySrc.size = %d\n, ReleaseCommand.sourceCellId.pLMNIdentity.size =%d\n,ReleaseCommand=%v\n",
			pLMNIdentitySrc.size,
			ReleaseCommand.sourceCellId.pLMNIdentity.size,
			ReleaseCommand)

		 */
	}

	index = 0
	HandoverCommandPtrs := (**C.HandoverCommand_t)(C.calloc(C.size_t(len(HandOverUEs)), 8))
	defer C.free(unsafe.Pointer(HandoverCommandPtrs))
	//var HandoverCommand_slice []*C.HandoverCommand_t
	for _,v := range HandOverUEs {
		var c_RNTI C.long
		var pLMNIdentitySrc C.OCTET_STRING_t
		var nRCellIdentitySrc C.BIT_STRING_t
		var pLMNIdentityDest C.OCTET_STRING_t
		var nRCellIdentityDest C.BIT_STRING_t

		//src
		c_RNTI = C.long(v.UEId)
		pLMNIdentitySrc.buf = (*C.uint8_t)(C.CBytes(GnoteB.Cells[v.SrcCellID].NRCGI.PlmnID.Buf))
		pLMNIdentitySrc.size = C.ulong(GnoteB.Cells[v.SrcCellID].NRCGI.PlmnID.Size)
		nRCellIdentitySrc.buf 			= (*C.uint8_t)(C.CBytes(GnoteB.Cells[v.SrcCellID].NRCGI.NRCellID.Buf))
		nRCellIdentitySrc.size 			= C.ulong(GnoteB.Cells[v.SrcCellID].NRCGI.NRCellID.Size)
		nRCellIdentitySrc.bits_unused 	= C.int(GnoteB.Cells[v.SrcCellID].NRCGI.NRCellID.BitsUnused)
		//defer C.free(unsafe.Pointer(pLMNIdentitySrc.buf))
		//defer C.free(unsafe.Pointer(nRCellIdentitySrc.buf))

		//dest
		pLMNIdentityDest.buf = (*C.uint8_t)(C.CBytes(GnoteB.Cells[v.DestCellID].NRCGI.PlmnID.Buf))
		pLMNIdentityDest.size = C.ulong(GnoteB.Cells[v.DestCellID].NRCGI.PlmnID.Size)
		nRCellIdentityDest.buf 			= (*C.uint8_t)(C.CBytes(GnoteB.Cells[v.DestCellID].NRCGI.NRCellID.Buf))
		nRCellIdentityDest.size 			= C.ulong(GnoteB.Cells[v.DestCellID].NRCGI.NRCellID.Size)
		nRCellIdentityDest.bits_unused 	= C.int(GnoteB.Cells[v.DestCellID].NRCGI.NRCellID.BitsUnused)
		//defer C.free(unsafe.Pointer(pLMNIdentityDest.buf))
		//defer C.free(unsafe.Pointer(nRCellIdentityDest.buf))

		HandoverCommand := C.e2sm_encode_ric_control_handover_command(pLMNIdentitySrc, nRCellIdentitySrc, c_RNTI,
			pLMNIdentityDest, nRCellIdentityDest)

		//HandoverCommand_slice = append(HandoverCommand_slice,HandoverCommand)
		*(**C.HandoverCommand_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(HandoverCommandPtrs)) + (uintptr)(8*C.int(index)))) = HandoverCommand

		//ReleaseCommand_slice = append(ReleaseCommand_slice,ReleaseCommand)
		index++

	}

	cptr := unsafe.Pointer(&buffer[0])
	size := C.e2sm_encode_ric_control_message_Format6(cptr, C.size_t(len(buffer)),
		ReleaseCommandPtrs,   C.int(len(ReleaseUEs)),
		HandoverCommandPtrs, C.int(len(HandOverUEs)))
	if size < 0 {
		return make([]byte, 0), errors.New("e2sm wrapper is unable to set ControlMessageFormat6 due to wrong or invalid input")
	}
	newBuffer = C.GoBytes(cptr, (C.int(size)+7)/8)
	return
}

func (c *E2sm) GetIndicationHeader(buffer []byte) (indHdr *IndicationHeader, err error) {
	cptr := unsafe.Pointer(&buffer[0])
	indHdr = &IndicationHeader{}
	decodedHdr := C.e2sm_decode_ric_indication_header(cptr, C.size_t(len(buffer)))
	if decodedHdr == nil {
		return indHdr, errors.New("e2sm wrapper is unable to get IndicationHeader due to wrong or invalid input")
	}
	defer C.e2sm_free_ric_indication_header(decodedHdr)
	return
}

func (c *E2sm) GetIndicationMessage(buffer []byte) (indMsg *IndicationMessage, err error) {
	cptr := unsafe.Pointer(&buffer[0])
	indMsg = &IndicationMessage{}
	decodedMsg := C.e2sm_decode_ric_indication_message(cptr, C.size_t(len(buffer)))
	if decodedMsg == nil {
		return indMsg, errors.New("e2sm wrapper is unable to get IndicationMessage due to wrong or invalid input")
	}
	defer C.e2sm_free_ric_indication_message(decodedMsg)

	indMsg.IndMsgType = int32(decodedMsg.rICIndicationMessageFormat.present)
	if indMsg.IndMsgType == 1 {
		indMsgFormat1 := &IndicationMessageFormat1{}
		indMsgFormat1_C := *(**C.E2SM_TS_IndicationMessageFormat1_t)(unsafe.Pointer(&decodedMsg.rICIndicationMessageFormat.choice[0]))
		indMsgFormat1.ContainerCount = int(indMsgFormat1_C.o_CU_CPCellLoadInfoContainer.o_CU_CPCellMeasurement_List.list.count)
		for i := 0; i < indMsgFormat1.ContainerCount; i++ {
			Container := &indMsgFormat1.o_Cu_CpCellMeasurement[i]
			var sizeof_O_Cu_CpCellMeasurement_t *C.O_Cu_CpCellMeasurement_t
			Container_C := *(**C.O_Cu_CpCellMeasurement_t)(unsafe.Pointer(uintptr(unsafe.Pointer(indMsgFormat1_C.o_CU_CPCellLoadInfoContainer.o_CU_CPCellMeasurement_List.list.array)) + (uintptr)(i)*unsafe.Sizeof(sizeof_O_Cu_CpCellMeasurement_t)))

			plmnID_C := Container_C.nRCgi.pLMNIdentity
			Container.NRCGI.PlmnID.Buf = C.GoBytes(unsafe.Pointer(plmnID_C.buf), C.int(plmnID_C.size))
			Container.NRCGI.PlmnID.Size = int(plmnID_C.size)
			//fmt.Printf("Container.NRCGI.PlmnID.Buf = %v\n,Container.NRCGI.PlmnID.Size = %d \n",Container.NRCGI.PlmnID.Buf,Container.NRCGI.PlmnID.Size)

			nRCellID := Container_C.nRCgi.nRCellIdentity
			Container.NRCGI.NRCellID.Buf = C.GoBytes(unsafe.Pointer(nRCellID.buf), C.int(nRCellID.size))
			Container.NRCGI.NRCellID.Size = int(nRCellID.size)
			Container.NRCGI.NRCellID.BitsUnused = int(nRCellID.bits_unused)
			CellId,err := c.ParseNRCGI(Container.NRCGI)
			if err != nil {
				return nil,err
			}
			Container.NRCellid = CellId
			Container.NumberOfSupportedRrcConnections = int(Container_C.numberOfSupportedRrcConnections)
			Container.NumberOfRrcConnections = int(Container_C.numberOfRrcConnections)
			/*
			fmt.Printf("CellId = %s,NumberOfSupportedRrcConnections = %d,\nNumberOfRrcConnections = %d\n",CellId,
				Container.NumberOfSupportedRrcConnections,
				Container.NumberOfRrcConnections)

			 */
		}
		indMsg.IndMsg = indMsgFormat1
	} else if (indMsg.IndMsgType == 4){
		indMsgFormat4 := &IndicationMessageFormat4{}
		indMsgFormat4_C := *(**C.E2SM_TS_IndicationMessageFormat4_t)(unsafe.Pointer(&decodedMsg.rICIndicationMessageFormat.choice[0]))
		indMsgFormat4.ContainerCount = int(indMsgFormat4_C.o_CU_CPUeMeasurementContainer.o_CU_CPUeMeasurement_List.list.count)
		for i := 0; i < indMsgFormat4.ContainerCount; i++ {
			Container := &indMsgFormat4.o_Cu_CpUeMeasurement[i]
			var sizeof_O_Cu_CpUeMeasurement_t *C.O_Cu_CpUeMeasurement_t
			Container_C := *(**C.O_Cu_CpUeMeasurement_t)(unsafe.Pointer(uintptr(unsafe.Pointer(indMsgFormat4_C.o_CU_CPUeMeasurementContainer.o_CU_CPUeMeasurement_List.list.array)) + (uintptr)(i)*unsafe.Sizeof(sizeof_O_Cu_CpUeMeasurement_t)))
			Container.UEId = int(*Container_C.uEId.c_RNTI)

			plmnID_C := Container_C.nRCgi.pLMNIdentity
			Container.NRCGI.PlmnID.Buf = C.GoBytes(unsafe.Pointer(plmnID_C.buf), C.int(plmnID_C.size))
			Container.NRCGI.PlmnID.Size = int(plmnID_C.size)

			nRCellID := Container_C.nRCgi.nRCellIdentity
			Container.NRCGI.NRCellID.Buf = C.GoBytes(unsafe.Pointer(nRCellID.buf), C.int(nRCellID.size))
			Container.NRCGI.NRCellID.Size = int(nRCellID.size)
			Container.NRCGI.NRCellID.BitsUnused = int(nRCellID.bits_unused)
			CellId,err := c.ParseNRCGI(Container.NRCGI)
			if err != nil {
				return nil,err
			}
			Container.NRCellid = CellId

			based_SSB := *(*C.MeasurementQuantityResult_t)(Container_C.measurementResult.measurementResultServingCell.based_SSB)
			Container.MResult.measurementResultServingCell.basedSSB.rSRP = int(based_SSB.rSRP)
			Container.MResult.measurementResultServingCell.basedSSB.rSRQ = int(based_SSB.rSRQ)
			Container.MResult.measurementResultServingCell.basedSSB.sINR = int(based_SSB.sINR)
/*
				fmt.Printf("CellId = %s,UEId = %d,rSRP = %d,rSRQ = %d,sINR = %d\n",CellId,Container.UEId,
				Container.MResult.measurementResultServingCell.basedSSB.rSRP,
				Container.MResult.measurementResultServingCell.basedSSB.rSRQ,
				Container.MResult.measurementResultServingCell.basedSSB.sINR)

 */


		}
		indMsg.IndMsg = indMsgFormat4
	}

	/*
	indMsg.StyleType = int64(decodedMsg.ric_Style_Type)

	indMsg.IndMsgType = int32(decodedMsg.indicationMessage.present)

	if indMsg.IndMsgType == 1 {
		indMsgFormat1 := &IndicationMessageFormat1{}
		indMsgFormat1_C := *(**C.E2SM_KPM_IndicationMessage_Format1_t)(unsafe.Pointer(&decodedMsg.indicationMessage.choice[0]))

		indMsgFormat1.PMContainerCount = int(indMsgFormat1_C.pm_Containers.list.count)
		for i := 0; i < indMsgFormat1.PMContainerCount; i++ {
			pmContainer := &indMsgFormat1.PMContainers[i]
			var sizeof_PM_Containers_List_t *C.PM_Containers_List_t
			pmContainer_C := *(**C.PM_Containers_List_t)(unsafe.Pointer(uintptr(unsafe.Pointer(indMsgFormat1_C.pm_Containers.list.array)) + (uintptr)(i)*unsafe.Sizeof(sizeof_PM_Containers_List_t)))

			if pmContainer_C.performanceContainer != nil {
				pfContainer := &PFContainerType{}

				pfContainer.ContainerType = int32(pmContainer_C.performanceContainer.present)

				if pfContainer.ContainerType == 1 {
					oDU_PF := &ODUPFContainerType{}
					oDU_PF_C := *(**C.ODU_PF_Container_t)(unsafe.Pointer(&pmContainer_C.performanceContainer.choice[0]))

					oDU_PF.CellResourceReportCount = int(oDU_PF_C.cellResourceReportList.list.count)
					for j := 0; j < oDU_PF.CellResourceReportCount; j++ {
						cellResourceReport := &oDU_PF.CellResourceReports[j]
						var sizeof_CellResourceReportListItem_t *C.CellResourceReportListItem_t
						cellResourceReport_C := *(**C.CellResourceReportListItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(oDU_PF_C.cellResourceReportList.list.array)) + (uintptr)(j)*unsafe.Sizeof(sizeof_CellResourceReportListItem_t)))

						cellResourceReport.NRCGI.PlmnID.Buf = C.GoBytes(unsafe.Pointer(cellResourceReport_C.nRCGI.pLMN_Identity.buf), C.int(cellResourceReport_C.nRCGI.pLMN_Identity.size))
						cellResourceReport.NRCGI.PlmnID.Size = int(cellResourceReport_C.nRCGI.pLMN_Identity.size)

						cellResourceReport.NRCGI.NRCellID.Buf = C.GoBytes(unsafe.Pointer(cellResourceReport_C.nRCGI.nRCellIdentity.buf), C.int(cellResourceReport_C.nRCGI.nRCellIdentity.size))
						cellResourceReport.NRCGI.NRCellID.Size = int(cellResourceReport_C.nRCGI.nRCellIdentity.size)
						cellResourceReport.NRCGI.NRCellID.BitsUnused = int(cellResourceReport_C.nRCGI.nRCellIdentity.bits_unused)

						if cellResourceReport_C.dl_TotalofAvailablePRBs != nil {
							cellResourceReport.TotalofAvailablePRBs.DL = int64(*cellResourceReport_C.dl_TotalofAvailablePRBs)
						} else {
							cellResourceReport.TotalofAvailablePRBs.DL = -1
						}

						if cellResourceReport_C.ul_TotalofAvailablePRBs != nil {
							cellResourceReport.TotalofAvailablePRBs.UL = int64(*cellResourceReport_C.ul_TotalofAvailablePRBs)
						} else {
							cellResourceReport.TotalofAvailablePRBs.UL = -1
						}

						cellResourceReport.ServedPlmnPerCellCount = int(cellResourceReport_C.servedPlmnPerCellList.list.count)
						for k := 0; k < cellResourceReport.ServedPlmnPerCellCount; k++ {
							servedPlmnPerCell := cellResourceReport.ServedPlmnPerCells[k]
							var sizeof_ServedPlmnPerCellListItem_t *C.ServedPlmnPerCellListItem_t
							servedPlmnPerCell_C := *(**C.ServedPlmnPerCellListItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(cellResourceReport_C.servedPlmnPerCellList.list.array)) + (uintptr)(k)*unsafe.Sizeof(sizeof_ServedPlmnPerCellListItem_t)))

							servedPlmnPerCell.PlmnID.Buf = C.GoBytes(unsafe.Pointer(servedPlmnPerCell_C.pLMN_Identity.buf), C.int(servedPlmnPerCell_C.pLMN_Identity.size))
							servedPlmnPerCell.PlmnID.Size = int(servedPlmnPerCell_C.pLMN_Identity.size)

							if servedPlmnPerCell_C.du_PM_5GC != nil {
								duPM5GC := &DUPM5GCContainerType{}
								duPM5GC_C := (*C.FGC_DU_PM_Container_t)(servedPlmnPerCell_C.du_PM_5GC)

								duPM5GC.SlicePerPlmnPerCellCount = int(duPM5GC_C.slicePerPlmnPerCellList.list.count)
								for l := 0; l < duPM5GC.SlicePerPlmnPerCellCount; l++ {
									slicePerPlmnPerCell := &duPM5GC.SlicePerPlmnPerCells[l]
									var sizeof_SlicePerPlmnPerCellListItem_t *C.SlicePerPlmnPerCellListItem_t
									slicePerPlmnPerCell_C := *(**C.SlicePerPlmnPerCellListItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(duPM5GC_C.slicePerPlmnPerCellList.list.array)) + (uintptr)(l)*unsafe.Sizeof(sizeof_SlicePerPlmnPerCellListItem_t)))

									slicePerPlmnPerCell.SliceID.SST.Buf = C.GoBytes(unsafe.Pointer(slicePerPlmnPerCell_C.sliceID.sST.buf), C.int(slicePerPlmnPerCell_C.sliceID.sST.size))
									slicePerPlmnPerCell.SliceID.SST.Size = int(slicePerPlmnPerCell_C.sliceID.sST.size)

									if slicePerPlmnPerCell_C.sliceID.sD != nil {
										slicePerPlmnPerCell.SliceID.SD = &OctetString{}
										slicePerPlmnPerCell.SliceID.SD.Buf = C.GoBytes(unsafe.Pointer(slicePerPlmnPerCell_C.sliceID.sD.buf), C.int(slicePerPlmnPerCell_C.sliceID.sD.size))
										slicePerPlmnPerCell.SliceID.SD.Size = int(slicePerPlmnPerCell_C.sliceID.sD.size)
									}

									slicePerPlmnPerCell.FQIPERSlicesPerPlmnPerCellCount = int(slicePerPlmnPerCell_C.fQIPERSlicesPerPlmnPerCellList.list.count)
									for m := 0; m < slicePerPlmnPerCell.FQIPERSlicesPerPlmnPerCellCount; m++ {
										fQIPerSlicesPerPlmnPerCell := &slicePerPlmnPerCell.FQIPERSlicesPerPlmnPerCells[m]
										var sizeof_FQIPERSlicesPerPlmnPerCellListItem_t *C.FQIPERSlicesPerPlmnPerCellListItem_t
										fQIPerSlicesPerPlmnPerCell_C := *(**C.FQIPERSlicesPerPlmnPerCellListItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(slicePerPlmnPerCell_C.fQIPERSlicesPerPlmnPerCellList.list.array)) + (uintptr)(m)*unsafe.Sizeof(sizeof_FQIPERSlicesPerPlmnPerCellListItem_t)))

										fQIPerSlicesPerPlmnPerCell.FiveQI = int64(fQIPerSlicesPerPlmnPerCell_C.fiveQI)

										if fQIPerSlicesPerPlmnPerCell_C.dl_PRBUsage != nil {
											fQIPerSlicesPerPlmnPerCell.PrbUsage.DL = int64(*fQIPerSlicesPerPlmnPerCell_C.dl_PRBUsage)
										} else {
											fQIPerSlicesPerPlmnPerCell.PrbUsage.DL = -1
										}

										if fQIPerSlicesPerPlmnPerCell_C.ul_PRBUsage != nil {
											fQIPerSlicesPerPlmnPerCell.PrbUsage.UL = int64(*fQIPerSlicesPerPlmnPerCell_C.ul_PRBUsage)
										} else {
											fQIPerSlicesPerPlmnPerCell.PrbUsage.UL = -1
										}
									}
								}

								servedPlmnPerCell.DUPM5GC = duPM5GC
							}

							if servedPlmnPerCell_C.du_PM_EPC != nil {
								duPMEPC := &DUPMEPCContainerType{}
								duPMEPC_C := (*C.EPC_DU_PM_Container_t)(servedPlmnPerCell_C.du_PM_EPC)

								duPMEPC.PerQCIReportCount = int(duPMEPC_C.perQCIReportList.list.count)
								for l := 0; l < duPMEPC.PerQCIReportCount; l++ {
									perQCIReport := &duPMEPC.PerQCIReports[l]
									var sizeof_PerQCIReportListItem_t *C.PerQCIReportListItem_t
									perQCIReport_C := *(**C.PerQCIReportListItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(duPMEPC_C.perQCIReportList.list.array)) + (uintptr)(l)*unsafe.Sizeof(sizeof_PerQCIReportListItem_t)))

									perQCIReport.QCI = int64(perQCIReport_C.qci)

									if perQCIReport_C.dl_PRBUsage != nil {
										perQCIReport.PrbUsage.DL = int64(*perQCIReport_C.dl_PRBUsage)
									} else {
										perQCIReport.PrbUsage.DL = -1
									}

									if perQCIReport_C.ul_PRBUsage != nil {
										perQCIReport.PrbUsage.UL = int64(*perQCIReport_C.ul_PRBUsage)
									} else {
										perQCIReport.PrbUsage.UL = -1
									}
								}

								servedPlmnPerCell.DUPMEPC = duPMEPC
							}
						}
					}

					pfContainer.Container = oDU_PF
				} else if pfContainer.ContainerType == 2 {
					oCU_CP_PF := &OCUCPPFContainerType{}
					oCU_CP_PF_C := *(**C.OCUCP_PF_Container_t)(unsafe.Pointer(&pmContainer_C.performanceContainer.choice[0]))

					if oCU_CP_PF_C.gNB_CU_CP_Name != nil {
						oCU_CP_PF.GNBCUCPName = &PrintableString{}
						oCU_CP_PF.GNBCUCPName.Buf = C.GoBytes(unsafe.Pointer(oCU_CP_PF_C.gNB_CU_CP_Name.buf), C.int(oCU_CP_PF_C.gNB_CU_CP_Name.size))
						oCU_CP_PF.GNBCUCPName.Size = int(oCU_CP_PF_C.gNB_CU_CP_Name.size)
					}

					if oCU_CP_PF_C.cu_CP_Resource_Status.numberOfActive_UEs != nil {
						oCU_CP_PF.CUCPResourceStatus.NumberOfActiveUEs = int64(*oCU_CP_PF_C.cu_CP_Resource_Status.numberOfActive_UEs)
					}

					pfContainer.Container = oCU_CP_PF
				} else if pfContainer.ContainerType == 3 {
					oCU_UP_PF := &OCUUPPFContainerType{}
					oCU_UP_PF_C := *(**C.OCUUP_PF_Container_t)(unsafe.Pointer(&pmContainer_C.performanceContainer.choice[0]))

					if oCU_UP_PF_C.gNB_CU_UP_Name != nil {
						oCU_UP_PF.GNBCUUPName = &PrintableString{}
						oCU_UP_PF.GNBCUUPName.Buf = C.GoBytes(unsafe.Pointer(oCU_UP_PF_C.gNB_CU_UP_Name.buf), C.int(oCU_UP_PF_C.gNB_CU_UP_Name.size))
						oCU_UP_PF.GNBCUUPName.Size = int(oCU_UP_PF_C.gNB_CU_UP_Name.size)
					}

					oCU_UP_PF.CUUPPFContainerItemCount = int(oCU_UP_PF_C.pf_ContainerList.list.count)
					for j := 0; j < oCU_UP_PF.CUUPPFContainerItemCount; j++ {
						cuUPPFContainer := &oCU_UP_PF.CUUPPFContainerItems[j]
						var sizeof_PF_ContainerListItem_t *C.PF_ContainerListItem_t
						cuUPPFContainer_C := *(**C.PF_ContainerListItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(oCU_UP_PF_C.pf_ContainerList.list.array)) + (uintptr)(j)*unsafe.Sizeof(sizeof_PF_ContainerListItem_t)))

						cuUPPFContainer.InterfaceType = int64(cuUPPFContainer_C.interface_type)

						cuUPPFContainer.OCUUPPMContainer.CUUPPlmnCount = int(cuUPPFContainer_C.o_CU_UP_PM_Container.plmnList.list.count)
						for k := 0; k < cuUPPFContainer.OCUUPPMContainer.CUUPPlmnCount; k++ {
							cuUPPlmn := &cuUPPFContainer.OCUUPPMContainer.CUUPPlmns[k]
							var sizeof_PlmnID_List_t *C.PlmnID_List_t
							cuUPPlmn_C := *(**C.PlmnID_List_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(cuUPPFContainer_C.o_CU_UP_PM_Container.plmnList.list.array)) + (uintptr)(k)*unsafe.Sizeof(sizeof_PlmnID_List_t)))

							cuUPPlmn.PlmnID.Buf = C.GoBytes(unsafe.Pointer(cuUPPlmn_C.pLMN_Identity.buf), C.int(cuUPPlmn_C.pLMN_Identity.size))
							cuUPPlmn.PlmnID.Size = int(cuUPPlmn_C.pLMN_Identity.size)

							if cuUPPlmn_C.cu_UP_PM_5GC != nil {
								cuUPPM5GC := &CUUPPM5GCType{}
								cuUPPM5GC_C := (*C.FGC_CUUP_PM_Format_t)(cuUPPlmn_C.cu_UP_PM_5GC)

								cuUPPM5GC.SliceToReportCount = int(cuUPPM5GC_C.sliceToReportList.list.count)
								for l := 0; l < cuUPPM5GC.SliceToReportCount; l++ {
									sliceToReport := &cuUPPM5GC.SliceToReports[l]
									var sizeof_SliceToReportListItem_t *C.SliceToReportListItem_t
									sliceToReport_C := *(**C.SliceToReportListItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(cuUPPM5GC_C.sliceToReportList.list.array)) + (uintptr)(l)*unsafe.Sizeof(sizeof_SliceToReportListItem_t)))

									sliceToReport.SliceID.SST.Buf = C.GoBytes(unsafe.Pointer(sliceToReport_C.sliceID.sST.buf), C.int(sliceToReport_C.sliceID.sST.size))
									sliceToReport.SliceID.SST.Size = int(sliceToReport_C.sliceID.sST.size)

									if sliceToReport_C.sliceID.sD != nil {
										sliceToReport.SliceID.SD = &OctetString{}
										sliceToReport.SliceID.SD.Buf = C.GoBytes(unsafe.Pointer(sliceToReport_C.sliceID.sD.buf), C.int(sliceToReport_C.sliceID.sD.size))
										sliceToReport.SliceID.SD.Size = int(sliceToReport_C.sliceID.sD.size)
									}

									sliceToReport.FQIPERSlicesPerPlmnCount = int(sliceToReport_C.fQIPERSlicesPerPlmnList.list.count)
									for m := 0; m < sliceToReport.FQIPERSlicesPerPlmnCount; m++ {
										fQIPerSlicesPerPlmn := &sliceToReport.FQIPERSlicesPerPlmns[m]
										var sizeof_FQIPERSlicesPerPlmnListItem_t *C.FQIPERSlicesPerPlmnListItem_t
										fQIPerSlicesPerPlmn_C := *(**C.FQIPERSlicesPerPlmnListItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(sliceToReport_C.fQIPERSlicesPerPlmnList.list.array)) + (uintptr)(m)*unsafe.Sizeof(sizeof_FQIPERSlicesPerPlmnListItem_t)))

										fQIPerSlicesPerPlmn.FiveQI = int64(fQIPerSlicesPerPlmn_C.fiveQI)

										if fQIPerSlicesPerPlmn_C.pDCPBytesDL != nil {
											fQIPerSlicesPerPlmn.PDCPBytesDL = &Integer{}
											fQIPerSlicesPerPlmn.PDCPBytesDL.Buf = C.GoBytes(unsafe.Pointer(fQIPerSlicesPerPlmn_C.pDCPBytesDL.buf), C.int(fQIPerSlicesPerPlmn_C.pDCPBytesDL.size))
											fQIPerSlicesPerPlmn.PDCPBytesDL.Size = int(fQIPerSlicesPerPlmn_C.pDCPBytesDL.size)
										}

										if fQIPerSlicesPerPlmn_C.pDCPBytesUL != nil {
											fQIPerSlicesPerPlmn.PDCPBytesUL = &Integer{}
											fQIPerSlicesPerPlmn.PDCPBytesUL.Buf = C.GoBytes(unsafe.Pointer(fQIPerSlicesPerPlmn_C.pDCPBytesUL.buf), C.int(fQIPerSlicesPerPlmn_C.pDCPBytesUL.size))
											fQIPerSlicesPerPlmn.PDCPBytesUL.Size = int(fQIPerSlicesPerPlmn_C.pDCPBytesUL.size)
										}
									}
								}

								cuUPPlmn.CUUPPM5GC = cuUPPM5GC
							}

							if cuUPPlmn_C.cu_UP_PM_EPC != nil {
								cuUPPMEPC := &CUUPPMEPCType{}
								cuUPPMEPC_C := (*C.EPC_CUUP_PM_Format_t)(cuUPPlmn_C.cu_UP_PM_EPC)

								cuUPPMEPC.CUUPPMEPCPerQCIReportCount = int(cuUPPMEPC_C.perQCIReportList.list.count)
								for l := 0; l < cuUPPMEPC.CUUPPMEPCPerQCIReportCount; l++ {
									perQCIReport := &cuUPPMEPC.CUUPPMEPCPerQCIReports[l]
									var sizeof_PerQCIReportListItemFormat_t *C.PerQCIReportListItemFormat_t
									perQCIReport_C := *(**C.PerQCIReportListItemFormat_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(cuUPPMEPC_C.perQCIReportList.list.array)) + (uintptr)(l)*unsafe.Sizeof(sizeof_PerQCIReportListItemFormat_t)))

									perQCIReport.QCI = int64(perQCIReport_C.qci)

									if perQCIReport_C.pDCPBytesDL != nil {
										perQCIReport.PDCPBytesDL = &Integer{}
										perQCIReport.PDCPBytesDL.Buf = C.GoBytes(unsafe.Pointer(perQCIReport_C.pDCPBytesDL.buf), C.int(perQCIReport_C.pDCPBytesDL.size))
										perQCIReport.PDCPBytesDL.Size = int(perQCIReport_C.pDCPBytesDL.size)
									}

									if perQCIReport_C.pDCPBytesUL != nil {
										perQCIReport.PDCPBytesUL = &Integer{}
										perQCIReport.PDCPBytesUL.Buf = C.GoBytes(unsafe.Pointer(perQCIReport_C.pDCPBytesUL.buf), C.int(perQCIReport_C.pDCPBytesUL.size))
										perQCIReport.PDCPBytesUL.Size = int(perQCIReport_C.pDCPBytesUL.size)
									}
								}

								cuUPPlmn.CUUPPMEPC = cuUPPMEPC
							}
						}
					}

					pfContainer.Container = oCU_UP_PF
				} else {
					return indMsg, errors.New("Unknown PF Container type")
				}

				pmContainer.PFContainer = pfContainer
			}

			if pmContainer_C.theRANContainer != nil {
				ranContainer := &RANContainerType{}

				ranContainer.Timestamp.Buf = C.GoBytes(unsafe.Pointer(pmContainer_C.theRANContainer.timestamp.buf), C.int(pmContainer_C.theRANContainer.timestamp.size))
				ranContainer.Timestamp.Size = int(pmContainer_C.theRANContainer.timestamp.size)

				ranContainer.ContainerType = int32(pmContainer_C.theRANContainer.reportContainer.present)

				if ranContainer.ContainerType == 1 {
					oDU_UE := &DUUsageReportType{}
					oDU_UE_C := *(**C.DU_Usage_Report_Per_UE_t)(unsafe.Pointer(&pmContainer_C.theRANContainer.reportContainer.choice[0]))

					oDU_UE.CellResourceReportItemCount = int(oDU_UE_C.cellResourceReportList.list.count)
					for j := 0; j < oDU_UE.CellResourceReportItemCount; j++ {
						cellResourceReport := &oDU_UE.CellResourceReportItems[j]
						var sizeof_DU_Usage_Report_CellResourceReportItem_t *C.DU_Usage_Report_CellResourceReportItem_t
						cellResourceReport_C := *(**C.DU_Usage_Report_CellResourceReportItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(oDU_UE_C.cellResourceReportList.list.array)) + (uintptr)(j)*unsafe.Sizeof(sizeof_DU_Usage_Report_CellResourceReportItem_t)))

						cellResourceReport.NRCGI.PlmnID.Buf = C.GoBytes(unsafe.Pointer(cellResourceReport_C.nRCGI.pLMN_Identity.buf), C.int(cellResourceReport_C.nRCGI.pLMN_Identity.size))
						cellResourceReport.NRCGI.PlmnID.Size = int(cellResourceReport_C.nRCGI.pLMN_Identity.size)

						cellResourceReport.NRCGI.NRCellID.Buf = C.GoBytes(unsafe.Pointer(cellResourceReport_C.nRCGI.nRCellIdentity.buf), C.int(cellResourceReport_C.nRCGI.nRCellIdentity.size))
						cellResourceReport.NRCGI.NRCellID.Size = int(cellResourceReport_C.nRCGI.nRCellIdentity.size)
						cellResourceReport.NRCGI.NRCellID.BitsUnused = int(cellResourceReport_C.nRCGI.nRCellIdentity.bits_unused)

						cellResourceReport.UeResourceReportItemCount = int(cellResourceReport_C.ueResourceReportList.list.count)
						for k := 0; k < cellResourceReport.UeResourceReportItemCount; k++ {
							ueResourceReport := &cellResourceReport.UeResourceReportItems[k]
							var sizeof_DU_Usage_Report_UeResourceReportItem_t *C.DU_Usage_Report_UeResourceReportItem_t
							ueResourceReport_C := *(**C.DU_Usage_Report_UeResourceReportItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(cellResourceReport_C.ueResourceReportList.list.array)) + (uintptr)(k)*unsafe.Sizeof(sizeof_DU_Usage_Report_UeResourceReportItem_t)))

							ueResourceReport.CRNTI.Buf = C.GoBytes(unsafe.Pointer(ueResourceReport_C.c_RNTI.buf), C.int(ueResourceReport_C.c_RNTI.size))
							ueResourceReport.CRNTI.Size = int(ueResourceReport_C.c_RNTI.size)

							if ueResourceReport_C.dl_PRBUsage != nil {
								ueResourceReport.PRBUsageDL = int64(*ueResourceReport_C.dl_PRBUsage)
							} else {
								ueResourceReport.PRBUsageDL = -1
							}

							if ueResourceReport_C.ul_PRBUsage != nil {
								ueResourceReport.PRBUsageUL = int64(*ueResourceReport_C.ul_PRBUsage)
							} else {
								ueResourceReport.PRBUsageUL = -1
							}
						}
					}

					ranContainer.Container = oDU_UE
				} else if ranContainer.ContainerType == 2 {
					oCU_CP_UE := &CUCPUsageReportType{}
					oCU_CP_UE_C := *(**C.CU_CP_Usage_Report_Per_UE_t)(unsafe.Pointer(&pmContainer_C.theRANContainer.reportContainer.choice[0]))

					oCU_CP_UE.CellResourceReportItemCount = int(oCU_CP_UE_C.cellResourceReportList.list.count)
					for j := 0; j < oCU_CP_UE.CellResourceReportItemCount; j++ {
						cellResourceReport := &oCU_CP_UE.CellResourceReportItems[j]
						var sizeof_CU_CP_Usage_Report_CellResourceReportItem_t *C.CU_CP_Usage_Report_CellResourceReportItem_t
						cellResourceReport_C := *(**C.CU_CP_Usage_Report_CellResourceReportItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(oCU_CP_UE_C.cellResourceReportList.list.array)) + (uintptr)(j)*unsafe.Sizeof(sizeof_CU_CP_Usage_Report_CellResourceReportItem_t)))

						cellResourceReport.NRCGI.PlmnID.Buf = C.GoBytes(unsafe.Pointer(cellResourceReport_C.nRCGI.pLMN_Identity.buf), C.int(cellResourceReport_C.nRCGI.pLMN_Identity.size))
						cellResourceReport.NRCGI.PlmnID.Size = int(cellResourceReport_C.nRCGI.pLMN_Identity.size)

						cellResourceReport.NRCGI.NRCellID.Buf = C.GoBytes(unsafe.Pointer(cellResourceReport_C.nRCGI.nRCellIdentity.buf), C.int(cellResourceReport_C.nRCGI.nRCellIdentity.size))
						cellResourceReport.NRCGI.NRCellID.Size = int(cellResourceReport_C.nRCGI.nRCellIdentity.size)
						cellResourceReport.NRCGI.NRCellID.BitsUnused = int(cellResourceReport_C.nRCGI.nRCellIdentity.bits_unused)

						cellResourceReport.UeResourceReportItemCount = int(cellResourceReport_C.ueResourceReportList.list.count)
						for k := 0; k < cellResourceReport.UeResourceReportItemCount; k++ {
							ueResourceReport := &cellResourceReport.UeResourceReportItems[k]
							var sizeof_CU_CP_Usage_Report_UeResourceReportItem_t *C.CU_CP_Usage_Report_UeResourceReportItem_t
							ueResourceReport_C := *(**C.CU_CP_Usage_Report_UeResourceReportItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(cellResourceReport_C.ueResourceReportList.list.array)) + (uintptr)(k)*unsafe.Sizeof(sizeof_CU_CP_Usage_Report_UeResourceReportItem_t)))

							ueResourceReport.CRNTI.Buf = C.GoBytes(unsafe.Pointer(ueResourceReport_C.c_RNTI.buf), C.int(ueResourceReport_C.c_RNTI.size))
							ueResourceReport.CRNTI.Size = int(ueResourceReport_C.c_RNTI.size)

							if ueResourceReport_C.serving_Cell_RF_Type != nil {
								ueResourceReport.ServingCellRF = &OctetString{}
								ueResourceReport.ServingCellRF.Buf = C.GoBytes(unsafe.Pointer(ueResourceReport_C.serving_Cell_RF_Type.buf), C.int(ueResourceReport_C.serving_Cell_RF_Type.size))
								ueResourceReport.ServingCellRF.Size = int(ueResourceReport_C.serving_Cell_RF_Type.size)
							}

							if ueResourceReport_C.neighbor_Cell_RF != nil {
								ueResourceReport.NeighborCellRF = &OctetString{}
								ueResourceReport.NeighborCellRF.Buf = C.GoBytes(unsafe.Pointer(ueResourceReport_C.neighbor_Cell_RF.buf), C.int(ueResourceReport_C.neighbor_Cell_RF.size))
								ueResourceReport.NeighborCellRF.Size = int(ueResourceReport_C.neighbor_Cell_RF.size)
							}
						}
					}

					ranContainer.Container = oCU_CP_UE
				} else if ranContainer.ContainerType == 3 {
					oCU_UP_UE := &CUUPUsageReportType{}
					oCU_UP_UE_C := *(**C.CU_UP_Usage_Report_Per_UE_t)(unsafe.Pointer(&pmContainer_C.theRANContainer.reportContainer.choice[0]))

					oCU_UP_UE.CellResourceReportItemCount = int(oCU_UP_UE_C.cellResourceReportList.list.count)
					for j := 0; j < oCU_UP_UE.CellResourceReportItemCount; j++ {
						cellResourceReport := &oCU_UP_UE.CellResourceReportItems[j]
						var sizeof_CU_UP_Usage_Report_CellResourceReportItem_t *C.CU_UP_Usage_Report_CellResourceReportItem_t
						cellResourceReport_C := *(**C.CU_UP_Usage_Report_CellResourceReportItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(oCU_UP_UE_C.cellResourceReportList.list.array)) + (uintptr)(j)*unsafe.Sizeof(sizeof_CU_UP_Usage_Report_CellResourceReportItem_t)))

						cellResourceReport.NRCGI.PlmnID.Buf = C.GoBytes(unsafe.Pointer(cellResourceReport_C.nRCGI.pLMN_Identity.buf), C.int(cellResourceReport_C.nRCGI.pLMN_Identity.size))
						cellResourceReport.NRCGI.PlmnID.Size = int(cellResourceReport_C.nRCGI.pLMN_Identity.size)

						cellResourceReport.NRCGI.NRCellID.Buf = C.GoBytes(unsafe.Pointer(cellResourceReport_C.nRCGI.nRCellIdentity.buf), C.int(cellResourceReport_C.nRCGI.nRCellIdentity.size))
						cellResourceReport.NRCGI.NRCellID.Size = int(cellResourceReport_C.nRCGI.nRCellIdentity.size)
						cellResourceReport.NRCGI.NRCellID.BitsUnused = int(cellResourceReport_C.nRCGI.nRCellIdentity.bits_unused)

						cellResourceReport.UeResourceReportItemCount = int(cellResourceReport_C.ueResourceReportList.list.count)
						for k := 0; k < cellResourceReport.UeResourceReportItemCount; k++ {
							ueResourceReport := &cellResourceReport.UeResourceReportItems[k]
							var sizeof_CU_UP_Usage_Report_UeResourceReportItem_t *C.CU_UP_Usage_Report_UeResourceReportItem_t
							ueResourceReport_C := *(**C.CU_UP_Usage_Report_UeResourceReportItem_t)(unsafe.Pointer((uintptr)(unsafe.Pointer(cellResourceReport_C.ueResourceReportList.list.array)) + (uintptr)(k)*unsafe.Sizeof(sizeof_CU_UP_Usage_Report_UeResourceReportItem_t)))

							ueResourceReport.CRNTI.Buf = C.GoBytes(unsafe.Pointer(ueResourceReport_C.c_RNTI.buf), C.int(ueResourceReport_C.c_RNTI.size))
							ueResourceReport.CRNTI.Size = int(ueResourceReport_C.c_RNTI.size)

							if ueResourceReport_C.pDCPBytesDL != nil {
								ueResourceReport.PDCPBytesDL = &Integer{}
								ueResourceReport.PDCPBytesDL.Buf = C.GoBytes(unsafe.Pointer(ueResourceReport_C.pDCPBytesDL.buf), C.int(ueResourceReport_C.pDCPBytesDL.size))
								ueResourceReport.PDCPBytesDL.Size = int(ueResourceReport_C.pDCPBytesDL.size)
							}

							if ueResourceReport_C.pDCPBytesUL != nil {
								ueResourceReport.PDCPBytesUL = &Integer{}
								ueResourceReport.PDCPBytesUL.Buf = C.GoBytes(unsafe.Pointer(ueResourceReport_C.pDCPBytesUL.buf), C.int(ueResourceReport_C.pDCPBytesUL.size))
								ueResourceReport.PDCPBytesUL.Size = int(ueResourceReport_C.pDCPBytesUL.size)
							}
						}
					}

					ranContainer.Container = oCU_UP_UE
				} else {
					return indMsg, errors.New("Unknown RAN Container type")
				}

				pmContainer.RANContainer = ranContainer
			}
		}

		indMsg.IndMsg = indMsgFormat1
	} else {
		return indMsg, errors.New("Unknown RIC Indication Message Format")
	}
*/
	return
}

func (c *E2sm) ParseNRCGI(nRCGI NRCGIType) (CellID string, err error) {
	var plmnID OctetString
	var nrCellID BitString

	plmnID = nRCGI.PlmnID
	CellID, _ = c.ParsePLMNIdentity(plmnID.Buf, plmnID.Size)

	nrCellID = nRCGI.NRCellID

	if plmnID.Size != 3 || nrCellID.Size != 5 {
		return "", errors.New("Invalid input: illegal length of NRCGI")
	}

	var former []uint8 = make([]uint8, 3)
	var latter []uint8 = make([]uint8, 6)

	former[0] = nrCellID.Buf[0] >> 4
	former[1] = nrCellID.Buf[0] & 0xf
	former[2] = nrCellID.Buf[1] >> 4
	latter[0] = nrCellID.Buf[1] & 0xf
	latter[1] = nrCellID.Buf[2] >> 4
	latter[2] = nrCellID.Buf[2] & 0xf
	latter[3] = nrCellID.Buf[3] >> 4
	latter[4] = nrCellID.Buf[3] & 0xf
	latter[5] = nrCellID.Buf[4] >> uint(nrCellID.BitsUnused)

	CellID = CellID + strconv.Itoa(int(former[0])) + strconv.Itoa(int(former[1])) + strconv.Itoa(int(former[2])) + strconv.Itoa(int(latter[0])) + strconv.Itoa(int(latter[1])) + strconv.Itoa(int(latter[2])) + strconv.Itoa(int(latter[3])) + strconv.Itoa(int(latter[4])) + strconv.Itoa(int(latter[5]))

	return
}

func (c *E2sm) ParsePLMNIdentity(buffer []byte, size int) (PlmnID string, err error) {
	if size != 3 {
		return "", errors.New("Invalid input: illegal length of PlmnID")
	}

	var mcc []uint8 = make([]uint8, 3)
	var mnc []uint8 = make([]uint8, 3)

	mcc[0] = buffer[0] >> 4
	mcc[1] = buffer[0] & 0xf
	mcc[2] = buffer[1] >> 4
	mnc[0] = buffer[1] & 0xf
	mnc[1] = buffer[2] >> 4
	mnc[2] = buffer[2] & 0xf

	if mnc[0] == 0xf {
		PlmnID = strconv.Itoa(int(mcc[0])) + strconv.Itoa(int(mcc[1])) + strconv.Itoa(int(mcc[2])) + strconv.Itoa(int(mnc[1])) + strconv.Itoa(int(mnc[2]))
	} else {
		PlmnID = strconv.Itoa(int(mcc[0])) + strconv.Itoa(int(mcc[1])) + strconv.Itoa(int(mcc[2])) + strconv.Itoa(int(mnc[0])) + strconv.Itoa(int(mnc[1])) + strconv.Itoa(int(mnc[2]))
	}

	return
}

func (c *E2sm) ParseSliceID(sliceID SliceIDType) (combined int32, err error) {
	if sliceID.SST.Size != 1 || (sliceID.SD != nil && sliceID.SD.Size != 3) {
		return 0, errors.New("Invalid input: illegal length of sliceID")
	}

	var temp uint8
	var sst int32
	var sd int32

	byteBuffer := bytes.NewBuffer(sliceID.SST.Buf)
	binary.Read(byteBuffer, binary.BigEndian, &temp)
	sst = int32(temp)

	if sliceID.SD == nil {
		combined = sst << 24
	} else {
		for i := 0; i < sliceID.SD.Size; i++ {
			byteBuffer = bytes.NewBuffer(sliceID.SD.Buf[i : i+1])
			binary.Read(byteBuffer, binary.BigEndian, &temp)
			sd = sd*256 + int32(temp)
		}
		combined = sst<<24 + sd
	}

	return
}

func (c *E2sm) ParseInteger(buffer []byte, size int) (value int64, err error) {
	var temp uint8
	var byteBuffer *bytes.Buffer

	for i := 0; i < size; i++ {
		byteBuffer = bytes.NewBuffer(buffer[i : i+1])
		binary.Read(byteBuffer, binary.BigEndian, &temp)
		value = value*256 + int64(temp)
	}

	return
}

func (c *E2sm) ParseTimestamp(buffer []byte, size int) (timestamp *Timestamp, err error) {
	var temp uint8
	var byteBuffer *bytes.Buffer
	var index int
	var sec int64
	var nsec int64

	for index := 0; index < size-8; index++ {
		byteBuffer = bytes.NewBuffer(buffer[index : index+1])
		binary.Read(byteBuffer, binary.BigEndian, &temp)
		sec = sec*256 + int64(temp)
	}

	for index = size - 8; index < size; index++ {
		byteBuffer = bytes.NewBuffer(buffer[index : index+1])
		binary.Read(byteBuffer, binary.BigEndian, &temp)
		nsec = nsec*256 + int64(temp)
	}

	timestamp = &Timestamp{TVsec: sec, TVnsec: nsec}
	return
}
/* cnbu
func (c *E2sm) SetHandoverCommand(buffer []byte, ReleaseUEs map[int]*UE, HandOverUEs map[int]*UE) (newBuffer []byte, err error) {
	cptr := unsafe.Pointer(&buffer[0])
	for i,UE := range ReleaseUEs {
		var (
			C_RNTI *C.long
			pLMNIdentitySrc C.OCTET_STRING_t
			nRCellIdentitySrc C.BIT_STRING_t

			pLMNIdentityDest C.OCTET_STRING_t
			nRCellIdentityDest C.BIT_STRING_t
		)
		C_RNTI = &i
		pLMNIdentitySrc.buf = []byte{UE.SrcCellID,}
	}

	size := C.e2sm_encode_ric_control_handover_command(cptr, C.size_t(len(buffer)), pLMNIdentitySrc, nRCellIdentitySrc, C_RNTI,
		 pLMNIdentityDest,  nRCellIdentityDest)
	if size < 0 {
		return make([]byte, 0), errors.New("e2sm wrapper is unable to set EventTriggerDefinition due to wrong or invalid input")
	}
	newBuffer = C.GoBytes(cptr, (C.int(size)+7)/8)
	return
}


func (c *E2sm) SetControlMessageFormat6(buffer []byte, ReleaseUEs map[int]*UE, HandOverUEs map[int]*UE) (newBuffer []byte, err error) {
	cptr := unsafe.Pointer(&buffer[0])
	var (
		HandoverCommand *C.HandoverCommand_t
		ReleaseCommand  *C.ReleaseCommand_t
		HandoverCommandSlice []byte
	)


	size := C.e2sm_encode_ric_control_message_Format6(cptr, C.size_t(len(buffer)), ReleaseCommand, len(ReleaseUEs), HandoverCommand,len(HandOverUEs))
	if size < 0 {
		return make([]byte, 0), errors.New("e2sm wrapper is unable to set ControlMessageFormat6 due to wrong or invalid input")
	}
	newBuffer = C.GoBytes(cptr, (C.int(size)+7)/8)
	return
}
*/