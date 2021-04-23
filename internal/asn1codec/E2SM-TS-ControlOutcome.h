/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_E2SM_TS_ControlOutcome_H_
#define	_E2SM_TS_ControlOutcome_H_


#include "asn_application.h"

/* Including external dependencies */
#include "RicControlOutcomeFormat-Choice.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* E2SM-TS-ControlOutcome */
typedef struct E2SM_TS_ControlOutcome {
	RicControlOutcomeFormat_Choice_t	 rICControlOutcomeFormat;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} E2SM_TS_ControlOutcome_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_E2SM_TS_ControlOutcome;

#ifdef __cplusplus
}
#endif

#endif	/* _E2SM_TS_ControlOutcome_H_ */
#include "asn_internal.h"
