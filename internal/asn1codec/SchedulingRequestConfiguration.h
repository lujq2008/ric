/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_SchedulingRequestConfiguration_H_
#define	_SchedulingRequestConfiguration_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum SchedulingRequestConfiguration__sr_ProhibitTimer {
	SchedulingRequestConfiguration__sr_ProhibitTimer_ms1	= 0,
	SchedulingRequestConfiguration__sr_ProhibitTimer_ms2	= 1,
	SchedulingRequestConfiguration__sr_ProhibitTimer_ms4	= 2,
	SchedulingRequestConfiguration__sr_ProhibitTimer_ms8	= 3,
	SchedulingRequestConfiguration__sr_ProhibitTimer_ms16	= 4,
	SchedulingRequestConfiguration__sr_ProhibitTimer_ms32	= 5,
	SchedulingRequestConfiguration__sr_ProhibitTimer_ms64	= 6,
	SchedulingRequestConfiguration__sr_ProhibitTimer_ms128	= 7
} e_SchedulingRequestConfiguration__sr_ProhibitTimer;
typedef enum SchedulingRequestConfiguration__sr_TransMax {
	SchedulingRequestConfiguration__sr_TransMax_n4	= 0,
	SchedulingRequestConfiguration__sr_TransMax_n8	= 1,
	SchedulingRequestConfiguration__sr_TransMax_n16	= 2,
	SchedulingRequestConfiguration__sr_TransMax_n32	= 3,
	SchedulingRequestConfiguration__sr_TransMax_n64	= 4,
	SchedulingRequestConfiguration__sr_TransMax_spare3	= 5,
	SchedulingRequestConfiguration__sr_TransMax_spare2	= 6,
	SchedulingRequestConfiguration__sr_TransMax_spare1	= 7
} e_SchedulingRequestConfiguration__sr_TransMax;

/* SchedulingRequestConfiguration */
typedef struct SchedulingRequestConfiguration {
	long	*sr_ProhibitTimer;	/* OPTIONAL */
	long	 sr_TransMax;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} SchedulingRequestConfiguration_t;

/* Implementation */
/* extern asn_TYPE_descriptor_t asn_DEF_sr_ProhibitTimer_2;	// (Use -fall-defs-global to expose) */
/* extern asn_TYPE_descriptor_t asn_DEF_sr_TransMax_11;	// (Use -fall-defs-global to expose) */
extern asn_TYPE_descriptor_t asn_DEF_SchedulingRequestConfiguration;
extern asn_SEQUENCE_specifics_t asn_SPC_SchedulingRequestConfiguration_specs_1;
extern asn_TYPE_member_t asn_MBR_SchedulingRequestConfiguration_1[2];

#ifdef __cplusplus
}
#endif

#endif	/* _SchedulingRequestConfiguration_H_ */
#include "asn_internal.h"
