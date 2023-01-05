import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateMembers } from "./types/archive/identity/tx";
import { MsgUpdateOperators } from "./types/archive/identity/tx";
import { MsgRegisterIssuer } from "./types/archive/identity/tx";
import { MsgAcceptIdentity } from "./types/archive/identity/tx";
import { MsgRenounceIdentity } from "./types/archive/identity/tx";
import { MsgRejectIdentity } from "./types/archive/identity/tx";
import { MsgIssueCertificate } from "./types/archive/identity/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/archive.identity.MsgUpdateMembers", MsgUpdateMembers],
    ["/archive.identity.MsgUpdateOperators", MsgUpdateOperators],
    ["/archive.identity.MsgRegisterIssuer", MsgRegisterIssuer],
    ["/archive.identity.MsgAcceptIdentity", MsgAcceptIdentity],
    ["/archive.identity.MsgRenounceIdentity", MsgRenounceIdentity],
    ["/archive.identity.MsgRejectIdentity", MsgRejectIdentity],
    ["/archive.identity.MsgIssueCertificate", MsgIssueCertificate],
    
];

export { msgTypes }