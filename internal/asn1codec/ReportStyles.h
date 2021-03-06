/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_ReportStyles_H_
#define	_ReportStyles_H_


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
struct RanParametersForReportAction;

/* ReportStyles */
typedef struct ReportStyles {
	RICStyleType_t	 rICReportStyleType;
	RICStyleName_t	 rICReportStyleName;
	RICFormatType_t	 rICReportActionFormatType;
	struct ReportStyles__rANParametersForReportAction_List {
		A_SEQUENCE_OF(struct RanParametersForReportAction) list;
		
		/* Context for parsing across buffer boundaries */
		asn_struct_ctx_t _asn_ctx;
	} rANParametersForReportAction_List;
	RICFormatType_t	 rICIndicationHeaderFormatType;
	RICFormatType_t	 rICIndicationMessageFormatType;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ReportStyles_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_ReportStyles;
extern asn_SEQUENCE_specifics_t asn_SPC_ReportStyles_specs_1;
extern asn_TYPE_member_t asn_MBR_ReportStyles_1[6];

#ifdef __cplusplus
}
#endif

#endif	/* _ReportStyles_H_ */
#include "asn_internal.h"
