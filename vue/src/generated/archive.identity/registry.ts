import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRenounceIdentity } from "./types/archive/identity/tx";
import { MsgRejectIdentity } from "./types/archive/identity/tx";
import { MsgAcceptIdentity } from "./types/archive/identity/tx";
import { MsgRegisterIssuer } from "./types/archive/identity/tx";
import { MsgIssueCertificate } from "./types/archive/identity/tx";
import { MsgRevokeIdentity } from "./types/archive/identity/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/archive.identity.MsgRenounceIdentity", MsgRenounceIdentity],
    ["/archive.identity.MsgRejectIdentity", MsgRejectIdentity],
    ["/archive.identity.MsgAcceptIdentity", MsgAcceptIdentity],
    ["/archive.identity.MsgRegisterIssuer", MsgRegisterIssuer],
    ["/archive.identity.MsgIssueCertificate", MsgIssueCertificate],
    ["/archive.identity.MsgRevokeIdentity", MsgRevokeIdentity],
    
];

export { msgTypes }