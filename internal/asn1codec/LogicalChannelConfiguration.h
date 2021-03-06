/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2SM-TS-IEs"
 * 	found in "E2SM-TS_ASN1_freeze_ver_0222.asn"
 * 	`asn1c -S /root/nric/asn1c/skeletons/ -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER`
 */

#ifndef	_LogicalChannelConfiguration_H_
#define	_LogicalChannelConfiguration_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeInteger.h"
#include "NativeEnumerated.h"
#include "constr_SEQUENCE.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum LogicalChannelConfiguration__prioritisedBitRate {
	LogicalChannelConfiguration__prioritisedBitRate_kBps0	= 0,
	LogicalChannelConfiguration__prioritisedBitRate_kBps8	= 1,
	LogicalChannelConfiguration__prioritisedBitRate_kBps16	= 2,
	LogicalChannelConfiguration__prioritisedBitRate_kBps32	= 3,
	LogicalChannelConfiguration__prioritisedBitRate_kBps64	= 4,
	LogicalChannelConfiguration__prioritisedBitRate_kBps128	= 5,
	LogicalChannelConfiguration__prioritisedBitRate_kBps256	= 6,
	LogicalChannelConfiguration__prioritisedBitRate_kBps512	= 7,
	LogicalChannelConfiguration__prioritisedBitRate_kBps1024	= 8,
	LogicalChannelConfiguration__prioritisedBitRate_kBps2048	= 9,
	LogicalChannelConfiguration__prioritisedBitRate_kBps4096	= 10,
	LogicalChannelConfiguration__prioritisedBitRate_kBps8192	= 11,
	LogicalChannelConfiguration__prioritisedBitRate_kBps16384	= 12,
	LogicalChannelConfiguration__prioritisedBitRate_kBps32768	= 13,
	LogicalChannelConfiguration__prioritisedBitRate_kBps65536	= 14,
	LogicalChannelConfiguration__prioritisedBitRate_infinity	= 15
} e_LogicalChannelConfiguration__prioritisedBitRate;
typedef enum LogicalChannelConfiguration__bucketSizeDuration {
	LogicalChannelConfiguration__bucketSizeDuration_ms5	= 0,
	LogicalChannelConfiguration__bucketSizeDuration_ms10	= 1,
	LogicalChannelConfiguration__bucketSizeDuration_ms20	= 2,
	LogicalChannelConfiguration__bucketSizeDuration_ms50	= 3,
	LogicalChannelConfiguration__bucketSizeDuration_ms100	= 4,
	LogicalChannelConfiguration__bucketSizeDuration_ms150	= 5,
	LogicalChannelConfiguration__bucketSizeDuration_ms300	= 6,
	LogicalChannelConfiguration__bucketSizeDuration_ms500	= 7,
	LogicalChannelConfiguration__bucketSizeDuration_ms1000	= 8,
	LogicalChannelConfiguration__bucketSizeDuration_spare7	= 9,
	LogicalChannelConfiguration__bucketSizeDuration_spare6	= 10,
	LogicalChannelConfiguration__bucketSizeDuration_spare5	= 11,
	LogicalChannelConfiguration__bucketSizeDuration_spare4	= 12,
	LogicalChannelConfiguration__bucketSizeDuration_spare3	= 13,
	LogicalChannelConfiguration__bucketSizeDuration_spare2	= 14,
	LogicalChannelConfiguration__bucketSizeDuration_spare1	= 15
} e_LogicalChannelConfiguration__bucketSizeDuration;

/* LogicalChannelConfiguration */
typedef struct LogicalChannelConfiguration {
	long	 logicalChannelGroup;
	long	 priority;
	long	 prioritisedBitRate;
	long	 bucketSizeDuration;
	/*
	 * This type is extensible,
	 * possible extensions are below.
	 */
	
	/* Context for parsing across buffer boundaries */
	asn_struct_ctx_t _asn_ctx;
} LogicalChannelConfiguration_t;

/* Implementation */
/* extern asn_TYPE_descriptor_t asn_DEF_prioritisedBitRate_4;	// (Use -fall-defs-global to expose) */
/* extern asn_TYPE_descriptor_t asn_DEF_bucketSizeDuration_21;	// (Use -fall-defs-global to expose) */
extern asn_TYPE_descriptor_t asn_DEF_LogicalChannelConfiguration;
extern asn_SEQUENCE_specifics_t asn_SPC_LogicalChannelConfiguration_specs_1;
extern asn_TYPE_member_t asn_MBR_LogicalChannelConfiguration_1[4];

#ifdef __cplusplus
}
#endif

#endif	/* _LogicalChannelConfiguration_H_ */
#include "asn_internal.h"
