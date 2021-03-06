/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_HandoverCommand_H_
#define	_HandoverCommand_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NRCGI.h"
#include "NativeInteger.h"
#include "UEID.h"
#include "NativeEnumerated.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum HandoverCommand__cause {
	HandoverCommand__cause_unspecified	= 0,
	HandoverCommand__cause_ricHandoverTriggered	= 1
} e_HandoverCommand__cause;

/* HandoverCommand */
typedef struct HandoverCommand {
	NRCGI_t	 sourceCellId;
	NRCGI_t	 targetCellId;
	long	*targetCellCarrierFrequency;	/* OPTIONAL */
	UEID_t	 uEId;
	long	*cause;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} HandoverCommand_t;

/* Implementation */
/* extern asn_TYPE_descriptor_t asn_DEF_cause_6;	// (Use -fall-defs-global to expose) */
extern asn_TYPE_descriptor_t asn_DEF_HandoverCommand;
extern asn_SEQUENCE_specifics_t asn_SPC_HandoverCommand_specs_1;
extern asn_TYPE_member_t asn_MBR_HandoverCommand_1[5];

#ifdef __cplusplus
}
#endif

#endif	/* _HandoverCommand_H_ */
#include "asn_internal.h"
