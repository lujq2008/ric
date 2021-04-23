/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_A1EventTriggerConfiguration_H_
#define	_A1EventTriggerConfiguration_H_


#include "asn_application.h"

/* Including external dependencies */
#include "BOOLEAN.h"
#include "NativeInteger.h"
#include "MeasurementTriggerQuantity.h"
#include "NativeEnumerated.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum A1EventTriggerConfiguration__timeToTrigger {
	A1EventTriggerConfiguration__timeToTrigger_ms0	= 0,
	A1EventTriggerConfiguration__timeToTrigger_ms40	= 1,
	A1EventTriggerConfiguration__timeToTrigger_ms64	= 2,
	A1EventTriggerConfiguration__timeToTrigger_ms80	= 3,
	A1EventTriggerConfiguration__timeToTrigger_ms100	= 4,
	A1EventTriggerConfiguration__timeToTrigger_ms128	= 5,
	A1EventTriggerConfiguration__timeToTrigger_ms160	= 6,
	A1EventTriggerConfiguration__timeToTrigger_ms256	= 7,
	A1EventTriggerConfiguration__timeToTrigger_ms320	= 8,
	A1EventTriggerConfiguration__timeToTrigger_ms480	= 9,
	A1EventTriggerConfiguration__timeToTrigger_ms512	= 10,
	A1EventTriggerConfiguration__timeToTrigger_ms640	= 11,
	A1EventTriggerConfiguration__timeToTrigger_ms1024	= 12,
	A1EventTriggerConfiguration__timeToTrigger_ms1280	= 13,
	A1EventTriggerConfiguration__timeToTrigger_ms2560	= 14,
	A1EventTriggerConfiguration__timeToTrigger_ms5120	= 15
} e_A1EventTriggerConfiguration__timeToTrigger;
typedef enum A1EventTriggerConfiguration__referenceSignalType {
	A1EventTriggerConfiguration__referenceSignalType_ssb	= 0,
	A1EventTriggerConfiguration__referenceSignalType_csi_rs	= 1
} e_A1EventTriggerConfiguration__referenceSignalType;
typedef enum A1EventTriggerConfiguration__reportInterval {
	A1EventTriggerConfiguration__reportInterval_ms120	= 0,
	A1EventTriggerConfiguration__reportInterval_ms240	= 1,
	A1EventTriggerConfiguration__reportInterval_ms480	= 2,
	A1EventTriggerConfiguration__reportInterval_ms640	= 3,
	A1EventTriggerConfiguration__reportInterval_ms1024	= 4,
	A1EventTriggerConfiguration__reportInterval_ms2048	= 5,
	A1EventTriggerConfiguration__reportInterval_ms5120	= 6,
	A1EventTriggerConfiguration__reportInterval_ms10240	= 7,
	A1EventTriggerConfiguration__reportInterval_ms20480	= 8,
	A1EventTriggerConfiguration__reportInterval_ms40960	= 9,
	A1EventTriggerConfiguration__reportInterval_min1	= 10,
	A1EventTriggerConfiguration__reportInterval_min6	= 11,
	A1EventTriggerConfiguration__reportInterval_min12	= 12,
	A1EventTriggerConfiguration__reportInterval_min30	= 13
} e_A1EventTriggerConfiguration__reportInterval;
typedef enum A1EventTriggerConfiguration__reportAmount {
	A1EventTriggerConfiguration__reportAmount_r1	= 0,
	A1EventTriggerConfiguration__reportAmount_r2	= 1,
	A1EventTriggerConfiguration__reportAmount_r4	= 2,
	A1EventTriggerConfiguration__reportAmount_r8	= 3,
	A1EventTriggerConfiguration__reportAmount_r16	= 4,
	A1EventTriggerConfiguration__reportAmount_r32	= 5,
	A1EventTriggerConfiguration__reportAmount_r64	= 6,
	A1EventTriggerConfiguration__reportAmount_infinity	= 7
} e_A1EventTriggerConfiguration__reportAmount;

/* Forward declarations */
struct MeasurementTriggerQuantity;

/* A1EventTriggerConfiguration */
typedef struct A1EventTriggerConfiguration {
	BOOLEAN_t	 enable;
	long	 index;
	MeasurementTriggerQuantity_t	 a1Threshold;
	BOOLEAN_t	 reportOnLeave;
	long	 hysteresis;
	long	 timeToTrigger;
	long	 referenceSignalType;
	long	 reportInterval;
	long	 reportAmount;
	MeasurementTriggerQuantity_t	 reportQuantityCell;
	long	 maxReportCells;
	struct MeasurementTriggerQuantity	*reportQuantityRsIndex;	/* OPTIONAL */
	BOOLEAN_t	 includeBeamMeasurements;
	long	*maxNrOfRsIndexesToReport;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} A1EventTriggerConfiguration_t;

/* Implementation */
/* extern asn_TYPE_descriptor_t asn_DEF_timeToTrigger_7;	// (Use -fall-defs-global to expose) */
/* extern asn_TYPE_descriptor_t asn_DEF_referenceSignalType_24;	// (Use -fall-defs-global to expose) */
/* extern asn_TYPE_descriptor_t asn_DEF_reportInterval_27;	// (Use -fall-defs-global to expose) */
/* extern asn_TYPE_descriptor_t asn_DEF_reportAmount_42;	// (Use -fall-defs-global to expose) */
extern asn_TYPE_descriptor_t asn_DEF_A1EventTriggerConfiguration;
extern asn_SEQUENCE_specifics_t asn_SPC_A1EventTriggerConfiguration_specs_1;
extern asn_TYPE_member_t asn_MBR_A1EventTriggerConfiguration_1[14];

#ifdef __cplusplus
}
#endif

#endif	/* _A1EventTriggerConfiguration_H_ */
#include "asn_internal.h"