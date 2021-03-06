/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_QuantityOffset_Choice_H_
#define	_QuantityOffset_Choice_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "constr_CHOICE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum QuantityOffset_Choice_PR {
	QuantityOffset_Choice_PR_NOTHING,	/* No components present */
	QuantityOffset_Choice_PR_rSRP,
	QuantityOffset_Choice_PR_rSRQ,
	QuantityOffset_Choice_PR_sINR
	/* Extensions may appear below */
	
} QuantityOffset_Choice_PR;

/* QuantityOffset-Choice */
typedef struct QuantityOffset_Choice {
	QuantityOffset_Choice_PR present;
	union QuantityOffset_Choice_u {
		long	 rSRP;
		long	 rSRQ;
		long	 sINR;
		/*
		 * This type is extensible,
		 * possible extensions are below.
		 */
	} choice;
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} QuantityOffset_Choice_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_QuantityOffset_Choice;
extern asn_CHOICE_specifics_t asn_SPC_QuantityOffset_Choice_specs_1;
extern asn_TYPE_member_t asn_MBR_QuantityOffset_Choice_1[3];
extern asn_per_constraints_t asn_PER_type_QuantityOffset_Choice_constr_1;

#ifdef __cplusplus
}
#endif

#endif	/* _QuantityOffset_Choice_H_ */
#include "asn_internal.h"
