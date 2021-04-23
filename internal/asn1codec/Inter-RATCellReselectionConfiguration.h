/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_Inter_RATCellReselectionConfiguration_H_
#define	_Inter_RATCellReselectionConfiguration_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "NativeEnumerated.h"
#include "BOOLEAN.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum Inter_RATCellReselectionConfiguration__allowedMeasurementBandwidth {
	Inter_RATCellReselectionConfiguration__allowedMeasurementBandwidth_mbw6	= 0,
	Inter_RATCellReselectionConfiguration__allowedMeasurementBandwidth_mbw15	= 1,
	Inter_RATCellReselectionConfiguration__allowedMeasurementBandwidth_mbw25	= 2,
	Inter_RATCellReselectionConfiguration__allowedMeasurementBandwidth_mbw50	= 3,
	Inter_RATCellReselectionConfiguration__allowedMeasurementBandwidth_mbw75	= 4,
	Inter_RATCellReselectionConfiguration__allowedMeasurementBandwidth_mbw100	= 5
} e_Inter_RATCellReselectionConfiguration__allowedMeasurementBandwidth;
typedef enum Inter_RATCellReselectionConfiguration__cellReselectionSubpriority {
	Inter_RATCellReselectionConfiguration__cellReselectionSubpriority_odot2	= 0,
	Inter_RATCellReselectionConfiguration__cellReselectionSubpriority_odot4	= 1,
	Inter_RATCellReselectionConfiguration__cellReselectionSubpriority_odot6	= 2,
	Inter_RATCellReselectionConfiguration__cellReselectionSubpriority_odot8	= 3
} e_Inter_RATCellReselectionConfiguration__cellReselectionSubpriority;

/* Inter-RATCellReselectionConfiguration */
typedef struct Inter_RATCellReselectionConfiguration {
	long	 dlCarrierFrequency;
	long	 allowedMeasurementBandwidth;
	BOOLEAN_t	 presenceAntennaPort1;
	long	 cellReselectionPriority;
	long	*cellReselectionSubpriority;	/* OPTIONAL */
	long	 threshX_HighP;
	long	 threshX_LowP;
	long	*threshX_HighQ;	/* OPTIONAL */
	long	*threshX_LowQ;	/* OPTIONAL */
	long	 qRxLevMin;
	long	*qQualMin;	/* OPTIONAL */
	long	*p_Max;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} Inter_RATCellReselectionConfiguration_t;

/* Implementation */
/* extern asn_TYPE_descriptor_t asn_DEF_allowedMeasurementBandwidth_3;	// (Use -fall-defs-global to expose) */
/* extern asn_TYPE_descriptor_t asn_DEF_cellReselectionSubpriority_12;	// (Use -fall-defs-global to expose) */
extern asn_TYPE_descriptor_t asn_DEF_Inter_RATCellReselectionConfiguration;
extern asn_SEQUENCE_specifics_t asn_SPC_Inter_RATCellReselectionConfiguration_specs_1;
extern asn_TYPE_member_t asn_MBR_Inter_RATCellReselectionConfiguration_1[12];

#ifdef __cplusplus
}
#endif

#endif	/* _Inter_RATCellReselectionConfiguration_H_ */
#include "asn_internal.h"
