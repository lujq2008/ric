/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_O_DuUeMeasurement_H_
#define	_O_DuUeMeasurement_H_


#include "asn_application.h"

/* Including external dependencies */
#include "UEID.h"
#include "NRCGI.h"
#include "NativeInteger.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* O-DuUeMeasurement */
typedef struct O_DuUeMeasurement {
	UEID_t	 uEId;
	NRCGI_t	 nRCgi;
	long	*uLUePrbUsedForDataTraffic;	/* OPTIONAL */
	long	*dLUePrbUsedForDataTraffic;	/* OPTIONAL */
	long	*averageDlUeThroughputInGnb;	/* OPTIONAL */
	long	*distributionOfDlUeThroughputInGnb;	/* OPTIONAL */
	long	*averageUlUeThroughputInGnb;	/* OPTIONAL */
	long	*uEMacRate;	/* OPTIONAL */
	long	*widebandCqiDistribution;	/* OPTIONAL */
	long	*averageMcs;	/* OPTIONAL */
	long	*tA;	/* OPTIONAL */
	long	*packetLossRate;	/* OPTIONAL */
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} O_DuUeMeasurement_t;

/* Implementation */
extern asn_TYPE_descriptor_t asn_DEF_O_DuUeMeasurement;
extern asn_SEQUENCE_specifics_t asn_SPC_O_DuUeMeasurement_specs_1;
extern asn_TYPE_member_t asn_MBR_O_DuUeMeasurement_1[12];

#ifdef __cplusplus
}
#endif

#endif	/* _O_DuUeMeasurement_H_ */
#include "asn_internal.h"