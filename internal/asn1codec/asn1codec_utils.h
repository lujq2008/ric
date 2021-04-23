
#ifndef INC_ASN1CODEC_UTILS_H_
#define INC_ASN1CODEC_UTILS_H_

#ifndef ASN_DISABLE_OER_SUPPORT
#define ASN_DISABLE_OER_SUPPORT
#endif

#ifndef ASN_PDU_COLLECTION
#define ASN_PDU_COLLECTION
#endif

#include <stdbool.h>
#include <E2AP-PDU.h>
#include <ProtocolIE-Field.h>
//#include "E2SM-TS-RANfunction-Description.h"
#include "E2SM-TS-RANFunctionDefinition.h"
//#include <ProtocolExtensionContainer.h>
//#include <ProtocolExtensionField.h>
#include <CriticalityDiagnostics-IE-List.h>

#define pLMN_Identity_size 3
#define shortMacro_eNB_ID_size  18
#define macro_eNB_ID_size       20
#define longMacro_eNB_ID_size   21
#define home_eNB_ID_size        28
#define eUTRANcellIdentifier_size 28

#ifdef __cplusplus
extern "C"
{
#endif

bool asn1_pdu_printer(const E2AP_PDU_t *pdu, size_t obufsz, char *buf);
bool asn1_ranfunction_printer(const E2SM_TS_RANFunctionDefinition_t* ranfunction, size_t obufsz, char* buf);
bool asn1_pdu_xer_printer(const E2AP_PDU_t *pdu, size_t obufsz, char *buf);
bool per_unpack_pdu(E2AP_PDU_t *pdu, size_t packed_buf_size, const void* packed_buf,size_t err_buf_size, char* err_buf);
bool per_pack_pdu(E2AP_PDU_t *pdu, size_t *packed_buf_size, unsigned char* packed_buf,size_t err_buf_size, char* err_buf);
bool unpack_pdu_aux(E2AP_PDU_t *pdu, size_t packed_buf_size, const void* packed_buf,size_t err_buf_size, char* err_buf,enum asn_transfer_syntax syntax);
bool pack_pdu_aux(E2AP_PDU_t *pdu, size_t *packed_buf_size, unsigned char* packed_buf,size_t err_buf_size, char* err_buf,enum asn_transfer_syntax syntax);
bool test_unpack(void);
bool per_unpack_ranfunction(E2SM_TS_RANFunctionDefinition_t* ranfunction, size_t packed_buf_size, const void* packed_buf, size_t err_buf_size, char* err_buf);
bool unpack_ranfunction_aux(E2SM_TS_RANFunctionDefinition_t* ranfunction, size_t packed_buf_size, const void* packed_buf, size_t err_buf_size, char* err_buf, enum asn_transfer_syntax syntax);

E2AP_PDU_t *new_pdu(size_t sz);
void delete_pdu(E2AP_PDU_t *pdu);

E2SM_TS_RANFunctionDefinition_t* new_ranfunction(size_t sz);
void delete_ranfunction(E2SM_TS_RANFunctionDefinition_t* ranfunction);

#ifdef __cplusplus
}
#endif

#endif /* INC_ASN1CODEC_UTILS_H_ */
