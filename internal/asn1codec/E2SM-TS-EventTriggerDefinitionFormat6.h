/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_E2SM_TS_EventTriggerDefinitionFormat6_H_
#define	_E2SM_TS_EventTriggerDefinitionFormat6_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* E2SM-TS-EventTriggerDefinitionFormat6 */
typedef struct E2SM_TS_EventTriggerDefinitionFormat6 {
	long	 dLPdcpSduDropRateInGnb_Cu_Up;
	long	*dLCellPdcpSduDataVolume;	/* OPTIONAL */
	long	*packetDelay;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} E2SM_TS_EventTriggerDefinitionFormat6_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_E2SM_TS_EventTriggerDefinitionFormat6;
extern asn_SEQUENCE_specifics_t asn_SPC_E2SM_TS_EventTriggerDefinitionFormat6_specs_1;
extern asn_TYPE_member_t asn_MBR_E2SM_TS_EventTriggerDefinitionFormat6_1[3];

#ifdef __cplusplus
}
#endif

#endif	/* _E2SM_TS_EventTriggerDefinitionFormat6_H_ */
#include "asn_internal.h"
