import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRenounceIdentity } from "./types/archive/identity/tx";
import { MsgUpdateOperators } from "./types/archive/identity/tx";
import { MsgRevokeIdentity } from "./types/archive/identity/tx";
import { MsgAcceptIdentity } from "./types/archive/identity/tx";
import { MsgRegisterIssuer } from "./types/archive/identity/tx";
import { MsgAddIdentityMember } from "./types/archive/identity/tx";
import { MsgIssueCertificate } from "./types/archive/identity/tx";
import { MsgRejectIdentity } from "./types/archive/identity/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/archive.identity.MsgRenounceIdentity", MsgRenounceIdentity],
    ["/archive.identity.MsgUpdateOperators", MsgUpdateOperators],
    ["/archive.identity.MsgRevokeIdentity", MsgRevokeIdentity],
    ["/archive.identity.MsgAcceptIdentity", MsgAcceptIdentity],
    ["/archive.identity.MsgRegisterIssuer", MsgRegisterIssuer],
    ["/archive.identity.MsgAddIdentityMember", MsgAddIdentityMember],
    ["/archive.identity.MsgIssueCertificate", MsgIssueCertificate],
    ["/archive.identity.MsgRejectIdentity", MsgRejectIdentity],
    
];

export { msgTypes }