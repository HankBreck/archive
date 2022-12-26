import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateCda } from "./types/archive/cda/tx";
import { MsgFinalizeCda } from "./types/archive/cda/tx";
import { MsgApproveCda } from "./types/archive/cda/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/archive.cda.MsgCreateCda", MsgCreateCda],
    ["/archive.cda.MsgFinalizeCda", MsgFinalizeCda],
    ["/archive.cda.MsgApproveCda", MsgApproveCda],
    
];

export { msgTypes }