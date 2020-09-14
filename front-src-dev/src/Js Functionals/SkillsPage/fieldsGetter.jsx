import {httpGet} from "../AlphaAPI";
import {getFieldsUrl} from "../urlNames";
import {defaultHeader} from "../defaultParamsAndHeaders";

let fieldsHolder;
export function getSkillFields(skillPageComponent) {
    fieldsHolder = skillPageComponent;
    const response = httpGet(getFieldsUrl, defaultHeader, handleSuccessGetSkillFields, handleDenyGetSkillFields);
}

export function handleSuccessGetSkillFields(value) {
    console.log("Success Get Skill Fields")
    console.log('server Value: ' + JSON.stringify(value))
    console.log('server data: ' + JSON.stringify(value.data))
    fieldsHolder ? fieldsHolder.setFields(value.data) : alert('ridiim')
}

export function handleDenyGetSkillFields(value) {
    console.log("Server Error: " + JSON.stringify(value))
}
