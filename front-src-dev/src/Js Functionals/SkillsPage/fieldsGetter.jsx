import {httpGet} from "../AlphaAPI";
import {getFieldsUrl} from "../urlNames";
import {defaultHeader} from "../defaultParamsAndHeaders";

export function getSkillFields() {
    const response = httpGet(getFieldsUrl, defaultHeader, handleSuccessGetSkillFields, handleDenyGetSkillFields);
}

export function handleSuccessGetSkillFields(value) {

}

export function handleDenyGetSkillFields(value) {

}
