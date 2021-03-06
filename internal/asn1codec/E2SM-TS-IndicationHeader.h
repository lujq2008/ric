/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_E2SM_TS_IndicationHeader_H_
#define	_E2SM_TS_IndicationHeader_H_


#include "asn_application.h"

/* Including external dependencies */
#include "RICStyleType.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct RicIndicationHeaderFormat_Choice;

/* E2SM-TS-IndicationHeader */
typedef struct E2SM_TS_IndicationHeader {
	RICStyleType_t	*rICStyleType;	/* OPTIONAL */
	struct RicIndicationHeaderFormat_Choice	*rICIndicationHeaderFormat;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} E2SM_TS_IndicationHeader_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_E2SM_TS_IndicationHeader;

#ifdef __cplusplus
}
#endif

#endif	/* _E2SM_TS_IndicationHeader_H_ */
#include "asn_internal.h"
