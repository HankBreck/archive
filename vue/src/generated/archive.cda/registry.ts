import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgApproveCda } from "./types/cda/tx";
import { MsgCreateCda } from "./types/cda/tx";
import { MsgFinalizeCda } from "./types/cda/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/archive.cda.MsgApproveCda", MsgApproveCda],
    ["/archive.cda.MsgCreateCda", MsgCreateCda],
    ["/archive.cda.MsgFinalizeCda", MsgFinalizeCda],
    
];

export { msgTypes }