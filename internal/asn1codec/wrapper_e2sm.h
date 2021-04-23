#ifndef	_WRAPPER_H_
#define	_WRAPPER_H_

#include "E2SM-TS-EventTriggerDefinition.h"
#include "E2SM-TS-EventTriggerDefinitionFormat1.h"
//#include "Trigger-ConditionIE-Item.h"
#include "E2SM-TS-ReportActionDefinition.h"
#include "E2SM-TS-ReportActionDefinitionFormat1.h"
#include "E2SM-TS-IndicationHeader.h"
#include "E2SM-TS-IndicationHeaderFormat1.h"
#include "E2SM-TS-ControlHeader.h"
#include "E2SM-TS-ControlMessage.h"
/*
#include "GlobalTSnode-ID.h"
#include "GlobalTSnode-gNB-ID.h"
#include "GlobalTSnode-en-gNB-ID.h"
#include "GlobalTSnode-ng-eNB-ID.h"
#include "GlobalTSnode-eNB-ID.h"
*/
#include "PLMN-Identity.h"
#include "GNB-ID-Choice.h"
#include "GNB-CU-UP-ID.h"
#include "GNB-DU-ID.h"
#include "ENGNB-ID.h"
#include "ENB-ID-Choice.h"
#include "ENB-ID.h"
//#include "NRCGI.h"
//#include "SNSSAI.h"
//#include "GNB-Name.h"
#include "E2SM-TS-IndicationMessage.h"
#include "E2SM-TS-IndicationMessageFormat1.h"
//#include "PM-Containers-List.h"
//#include "PF-Container.h"
//#include "RAN-Container.h"
//#include "ODU-PF-Container.h"
/*
#include "CellResourceReportListItem.h"
#include "ServedPlmnPerCellListItem.h"
#include "FGC-DU-PM-Container.h"
#include "EPC-DU-PM-Container.h"
#include "SlicePerPlmnPerCellListItem.h"
#include "FQIPERSlicesPerPlmnPerCellListItem.h"
#include "PerQCIReportListItem.h"
#include "OCUCP-PF-Container.h"
#include "OCUUP-PF-Container.h"
#include "PF-ContainerListItem.h"
#include "PlmnID-List.h"
#include "FGC-CUUP-PM-Format.h"
#include "SliceToReportListItem.h"
#include "FQIPERSlicesPerPlmnListItem.h"
#include "EPC-CUUP-PM-Format.h"
#include "PerQCIReportListItemFormat.h"
#include "DU-Usage-Report-Per-UE.h"
#include "DU-Usage-Report-CellResourceReportItem.h"
#include "DU-Usage-Report-UeResourceReportItem.h"
#include "CU-CP-Usage-Report-Per-UE.h"
#include "CU-CP-Usage-Report-CellResourceReportItem.h"
#include "CU-CP-Usage-Report-UeResourceReportItem.h"
#include "CU-UP-Usage-Report-Per-UE.h"
#include "CU-UP-Usage-Report-CellResourceReportItem.h"
#include "CU-UP-Usage-Report-UeResourceReportItem.h"
*/
#include "HandoverCommand.h"
#include "ReleaseCommand.h"
#include "E2SM-TS-ControlMessageFormat6.h"
#include "UeHandover.h"
#include "UeRelease.h"
#include "E2SM-TS-IndicationMessageFormat4.h"
#include "O-Cu-CpUeMeasurement.h"
#include "MeasurementQuantityResult.h"
#include "O-Cu-CpCellMeasurement.h"
#include "MeasurementResultNR.h"


//int CommandSetLen = 256;
ssize_t e2sm_encode_ric_event_trigger_definition(void *buffer, size_t buf_size, size_t event_trigger_count, long *RT_periods);
ssize_t e2sm_encode_ric_action_definition(void *buffer, size_t buf_size, long ric_style_type);
E2SM_TS_IndicationHeader_t* e2sm_decode_ric_indication_header(void *buffer, size_t buf_size);
void e2sm_free_ric_indication_header(E2SM_TS_IndicationHeader_t* indHdr);
E2SM_TS_IndicationMessage_t* e2sm_decode_ric_indication_message(void *buffer, size_t buf_size);
void e2sm_free_ric_indication_message(E2SM_TS_IndicationMessage_t* indMsg);
ssize_t e2sm_encode_ric_control_message_Format6(void* buffer, size_t buf_size,
	ReleaseCommand_t** ReleaseCommand, int ReleaseCommandLen, HandoverCommand_t** HandoverCommand, int HandoverCommandLen);
HandoverCommand_t* e2sm_encode_ric_control_handover_command(OCTET_STRING_t pLMNIdentitySrc, BIT_STRING_t nRCellIdentitySrc, long c_RNTI,
	OCTET_STRING_t pLMNIdentityDest, BIT_STRING_t nRCellIdentityDest);
ReleaseCommand_t* e2sm_encode_ric_control_release_command(OCTET_STRING_t pLMNIdentitySrc, BIT_STRING_t nRCellIdentitySrc, long c_RNTI);
ssize_t e2sm_encode_ric_control_header(void* buffer, size_t buf_size, long ric_style_type);
#endif /* _WRAPPER_H_ */
