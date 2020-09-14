import {httpExcGET} from "../AlphaAPI";
import {addSkillUrl} from "../urlNames";
import {defaultHeader} from "../defaultParamsAndHeaders";


export function addSkillPost(skillName,mageType){
    let data = {
        "name" : skillName
    }
    mageType === "post" ? httpExcGET(mageType, addSkillUrl,data,addSkillSuccess,serverDeny,defaultHeader) : httpExcGET(mageType,data, addSkillUrl,removeSkillSuccess,serverDeny,defaultHeader)
}

export function addSkillSuccess(){
    console.log("Skill added.")
}

export function serverDeny(){
    console.log("Server Error: " + JSON.stringify(value))
}

export function removeSkillSuccess(){
    console.log("Skill removed.")
}
