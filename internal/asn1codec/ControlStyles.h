/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_ControlStyles_H_
#define	_ControlStyles_H_


#include "asn_application.h"

/* Including external dependencies */
#include "RICStyleType.h"
#include "RICStyleName.h"
#include "RICFormatType.h"
#include "asn_SEQUENCE_OF.h"
#include "constr_SEQUENCE_OF.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct RanParametersForControlMessage;
struct RanParametersForControlOutcome;

/* ControlStyles */
typedef struct ControlStyles {
	RICStyleType_t	 rICControlStyleType;
	RICStyleName_t	 rICControlStyleName;
	RICFormatType_t	 rICControlHeaderFormatType;
	RICFormatType_t	 rICControlMessageFormatType;
	struct ControlStyles__rANParametersForControlMessage_List {
		A_SEQUENCE_OF(struct RanParametersForControlMessage) list;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} rANParametersForControlMessage_List;
	RICFormatType_t	 rICControlOutcomeFormatType;
	struct ControlStyles__rANParametersForControlOutcome_List {
		A_SEQUENCE_OF(struct RanParametersForControlOutcome) list;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} rANParametersForControlOutcome_List;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ControlStyles_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_ControlStyles;
extern asn_SEQUENCE_specifics_t asn_SPC_ControlStyles_specs_1;
extern asn_TYPE_member_t asn_MBR_ControlStyles_1[7];

#ifdef __cplusplus
}
#endif

#endif	/* _ControlStyles_H_ */
#include "asn_internal.h"
