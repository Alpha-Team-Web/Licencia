import {httpExcGET} from "../AlphaAPI";
import {skillManagementUrl} from "../urlNames";
import {defaultHeader} from "../defaultParamsAndHeaders";


export function skillManagementPost(skillName, manageType) {
    let data = {
        "name": skillName
    }
    httpExcGET(manageType, skillManagementUrl, data,
        manageType === "post" ? addSkillSuccess : removeSkillSuccess, serverDeny,
        defaultHeader)
}

function addSkillSuccess() {
    console.log("Skill added.")
}

function serverDeny(value) {
    console.log("Server Error: " + JSON.stringify(value))
}

function removeSkillSuccess() {
    console.log("Skill removed.")
}
