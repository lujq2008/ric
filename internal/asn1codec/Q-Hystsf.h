/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_Q_Hystsf_H_
#define	_Q_Hystsf_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum Q_Hystsf__sf_Medium {
	Q_Hystsf__sf_Medium_dB_6	= 0,
	Q_Hystsf__sf_Medium_dB_4	= 1,
	Q_Hystsf__sf_Medium_dB_2	= 2,
	Q_Hystsf__sf_Medium_dB0	= 3
} e_Q_Hystsf__sf_Medium;
typedef enum Q_Hystsf__sf_High {
	Q_Hystsf__sf_High_dB_6	= 0,
	Q_Hystsf__sf_High_dB_4	= 1,
	Q_Hystsf__sf_High_dB_2	= 2,
	Q_Hystsf__sf_High_dB0	= 3
} e_Q_Hystsf__sf_High;

/* Q-Hystsf */
typedef struct Q_Hystsf {
	long	*sf_Medium;	/* OPTIONAL */
	long	*sf_High;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} Q_Hystsf_t;

/* Implementation */
/* extern asn_TYPE_descriptor_t asn_DEF_sf_Medium_2;	// (Use -fall-defs-global to expose) */
/* extern asn_TYPE_descriptor_t asn_DEF_sf_High_7;	// (Use -fall-defs-global to expose) */
extern asn_TYPE_descriptor_t asn_DEF_Q_Hystsf;
extern asn_SEQUENCE_specifics_t asn_SPC_Q_Hystsf_specs_1;
extern asn_TYPE_member_t asn_MBR_Q_Hystsf_1[2];

#ifdef __cplusplus
}
#endif

#endif	/* _Q_Hystsf_H_ */
#include "asn_internal.h"