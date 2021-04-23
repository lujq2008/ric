/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_ReportCellsByControl_H_
#define	_ReportCellsByControl_H_


#include "asn_application.h"

/* Including external dependencies */
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Forward declarations */
struct NRCGI;

/* ReportCellsByControl */
typedef struct ReportCellsByControl {
	struct NRCGI	*nRCgi;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} ReportCellsByControl_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_ReportCellsByControl;
extern asn_SEQUENCE_specifics_t asn_SPC_ReportCellsByControl_specs_1;
extern asn_TYPE_member_t asn_MBR_ReportCellsByControl_1[1];

#ifdef __cplusplus
}
#endif

#endif	/* _ReportCellsByControl_H_ */
#include "asn_internal.h"
