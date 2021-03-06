/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_UEID_H_
#define	_UEID_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "INTEGER.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* UEID */
typedef struct UEID {
	unsigned long	*rANUeNgapId;	/* OPTIONAL */
	INTEGER_t	*aMFUeNgapId;	/* OPTIONAL */
	unsigned long	*gNB_DUUeF1ApId;	/* OPTIONAL */
	unsigned long	*gNB_CUUeF1ApId;	/* OPTIONAL */
	unsigned long	*gNB_CU_CPUeE1ApId;	/* OPTIONAL */
	unsigned long	*gNB_CU_UPUeE1ApId;	/* OPTIONAL */
	unsigned long	*sourceNg_RanNodeUeXnapId;	/* OPTIONAL */
	unsigned long	*targetNg_RanNodeUeXnapId;	/* OPTIONAL */
	long	*c_RNTI;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} UEID_t;

/* Implementation */
/* extern asn_TYPE_descriptor_t asn_DEF_rANUeNgapId_2;	// (Use -fall-defs-global to expose) */
/* extern asn_TYPE_descriptor_t asn_DEF_gNB_DUUeF1ApId_4;	// (Use -fall-defs-global to expose) */
/* extern asn_TYPE_descriptor_t asn_DEF_gNB_CUUeF1ApId_5;	// (Use -fall-defs-global to expose) */
/* extern asn_TYPE_descriptor_t asn_DEF_gNB_CU_CPUeE1ApId_6;	// (Use -fall-defs-global to expose) */
/* extern asn_TYPE_descriptor_t asn_DEF_gNB_CU_UPUeE1ApId_7;	// (Use -fall-defs-global to expose) */
/* extern asn_TYPE_descriptor_t asn_DEF_sourceNg_RanNodeUeXnapId_8;	// (Use -fall-defs-global to expose) */
/* extern asn_TYPE_descriptor_t asn_DEF_targetNg_RanNodeUeXnapId_9;	// (Use -fall-defs-global to expose) */
extern asn_TYPE_descriptor_t asn_DEF_UEID;
extern asn_SEQUENCE_specifics_t asn_SPC_UEID_specs_1;
extern asn_TYPE_member_t asn_MBR_UEID_1[9];

#ifdef __cplusplus
}
#endif

#endif	/* _UEID_H_ */
#include "asn_internal.h"
