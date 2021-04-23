/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_EutraCellContainers_H_
#define	_EutraCellContainers_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "OCTET_STRING.h"
#include "NativeEnumerated.h"
#include "BOOLEAN.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum EutraCellContainers__qOffset {
	EutraCellContainers__qOffset_dB_24	= 0,
	EutraCellContainers__qOffset_dB_22	= 1,
	EutraCellContainers__qOffset_dB_20	= 2,
	EutraCellContainers__qOffset_dB_18	= 3,
	EutraCellContainers__qOffset_dB_16	= 4,
	EutraCellContainers__qOffset_dB_14	= 5,
	EutraCellContainers__qOffset_dB_12	= 6,
	EutraCellContainers__qOffset_dB_10	= 7,
	EutraCellContainers__qOffset_dB_8	= 8,
	EutraCellContainers__qOffset_dB_6	= 9,
	EutraCellContainers__qOffset_dB_5	= 10,
	EutraCellContainers__qOffset_dB_4	= 11,
	EutraCellContainers__qOffset_dB_3	= 12,
	EutraCellContainers__qOffset_dB_2	= 13,
	EutraCellContainers__qOffset_dB_1	= 14,
	EutraCellContainers__qOffset_dB0	= 15,
	EutraCellContainers__qOffset_dB1	= 16,
	EutraCellContainers__qOffset_dB2	= 17,
	EutraCellContainers__qOffset_dB3	= 18,
	EutraCellContainers__qOffset_dB4	= 19,
	EutraCellContainers__qOffset_dB5	= 20,
	EutraCellContainers__qOffset_dB6	= 21,
	EutraCellContainers__qOffset_dB8	= 22,
	EutraCellContainers__qOffset_dB10	= 23,
	EutraCellContainers__qOffset_dB12	= 24,
	EutraCellContainers__qOffset_dB14	= 25,
	EutraCellContainers__qOffset_dB16	= 26,
	EutraCellContainers__qOffset_dB18	= 27,
	EutraCellContainers__qOffset_dB20	= 28,
	EutraCellContainers__qOffset_dB22	= 29,
	EutraCellContainers__qOffset_dB24	= 30
} e_EutraCellContainers__qOffset;

/* Forward declarations */
struct E_UTRACGI;

/* EutraCellContainers */
typedef struct EutraCellContainers {
	struct E_UTRACGI	*eCGI;	/* OPTIONAL */
	long	*pCI;	/* OPTIONAL */
	OCTET_STRING_t	*tAC;	/* OPTIONAL */
	long	*eUTRACarrierArfcn;	/* OPTIONAL */
	long	*qOffset;	/* OPTIONAL */
	long	*qRxLevMinOffsetCell;	/* OPTIONAL */
	long	*qQualLevMinOffsetCell;	/* OPTIONAL */
	BOOLEAN_t	*isBlackCell;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} EutraCellContainers_t;

/* Implementation */
/* extern asn_TYPE_descriptor_t asn_DEF_qOffset_6;	// (Use -fall-defs-global to expose) */
extern asn_TYPE_descriptor_t asn_DEF_EutraCellContainers;
extern asn_SEQUENCE_specifics_t asn_SPC_EutraCellContainers_specs_1;
extern asn_TYPE_member_t asn_MBR_EutraCellContainers_1[8];

#ifdef __cplusplus
}
#endif

#endif	/* _EutraCellContainers_H_ */
#include "asn_internal.h"
