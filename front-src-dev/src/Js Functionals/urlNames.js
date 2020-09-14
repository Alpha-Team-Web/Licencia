
export const serverAddress = "http://localhost:8008/";

export const urlSignUp = serverAddress + "io/register"
export const urlLogin = serverAddress + "io/login"

// ProfilePart

export const urlGetEmployerProfileInfo = serverAddress + "employer/profile/get"
export const urlGetFreelancerProfileInfo = serverAddress + "freelancer/profile/get"
export const saveProfileUrlEmployer = serverAddress + "employer/profile/general";
export const changePasswordUrlEmployer = serverAddress + "employer/profile/password";
export const saveProfileUrlFreeLancer = serverAddress + "freelancer/profile/general";
export const saveGithubUrlFreeLancer = serverAddress + "freelancer/profile/links";
export const changePasswordUrlFreeLancer = serverAddress + "freelancer/profile/password";
export const uploadProfilePicUrlEmployer = serverAddress + "files/profile-pic/employer/upload";
export const uploadProfilePicUrlFreelancer = serverAddress + "files/profile-pic/freelancer/upload";

// ProfilePart


// SkillsPart

export const getFieldsUrl = serverAddress + "/ProfilePart";
export const getSkillsByFieldIdUrl = serverAddress + "/fields/field-skills";
export const skillManagementUrl = serverAddress +  "/freelancer/profile/skills";

// External WebSites
export const gitHubUrl = "https://github.com/"


// --disable-web-security --user-data-dir="[some directory here]"


//Server Errors: view notifications//notif.go
