/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_UeFailedToHandover_H_
#define	_UeFailedToHandover_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NRCGI.h"
#include "UEID.h"
#include "NativeEnumerated.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum UeFailedToHandover__cause {
	UeFailedToHandover__cause_unspecified	= 0,
	UeFailedToHandover__cause_unknowntargetUE	= 1,
	UeFailedToHandover__cause_handovertargetnotallowed	= 2
} e_UeFailedToHandover__cause;

/* UeFailedToHandover */
typedef struct UeFailedToHandover {
	NRCGI_t	 sourceNrCgi;
	NRCGI_t	 targetNrCgi;
	UEID_t	 uEId;
	long	 cause;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} UeFailedToHandover_t;

/* Implementation */
/* extern asn_TYPE_descriptor_t asn_DEF_cause_5;	// (Use -fall-defs-global to expose) */
extern asn_TYPE_descriptor_t asn_DEF_UeFailedToHandover;
extern asn_SEQUENCE_specifics_t asn_SPC_UeFailedToHandover_specs_1;
extern asn_TYPE_member_t asn_MBR_UeFailedToHandover_1[4];

#ifdef __cplusplus
}
#endif

#endif	/* _UeFailedToHandover_H_ */
#include "asn_internal.h"
