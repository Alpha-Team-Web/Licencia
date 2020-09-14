import {httpGet} from "../AlphaAPI";
import {getFieldsUrl} from "../urlNames";
import {defaultHeader} from "../defaultParamsAndHeaders";

export function getSkillFields(fieldsSetter) {
    const response = httpGet(getFieldsUrl, defaultHeader,
        (value) => handleSuccessGetSkillFields(fieldsSetter, value),
        handleDenyGetSkillFields);
}

export function handleSuccessGetSkillFields(setFields, value) {
    console.log("Success Get Skill Fields")
    console.log('server Value: ' + JSON.stringify(value))
    console.log('server data: ' + JSON.stringify(value.data))
    setFields ? setFields(value.data) : alert('ridiim')
}

export function handleDenyGetSkillFields(value) {
    console.log("Server Error: " + JSON.stringify(value))
}
