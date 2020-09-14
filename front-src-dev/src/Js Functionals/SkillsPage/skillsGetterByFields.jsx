import {httpGet} from "../AlphaAPI";
import {getSkillsByFieldIdUrl} from "../urlNames";
import {defaultHeader, getParam} from "../defaultParamsAndHeaders";

export function getSkillsByFieldId(skillsSetter, fieldId) {
    const response = httpGet(getSkillsByFieldIdUrl, defaultHeader,
        (value) => successGetSkillsByFieldId(skillsSetter, value),
        denyGetSkillsByFieldId, getParam('field-id', fieldId))
}

function successGetSkillsByFieldId(setSkills, value) {
    console.log("Success Get Skill Fields")
    console.log('server Value: ' + JSON.stringify(value))
    console.log('server data: ' + JSON.stringify(value.data))
    setSkills ? setSkills(value.data) : alert('ridiim')
}

function denyGetSkillsByFieldId(value) {
    console.log("Server Error: " + JSON.stringify(value))
}
