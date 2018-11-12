import {BaseURL, axios} from '../config'
import qs from 'qs'

// get project list
export const getProjectList = params => { return axios.get(`${BaseURL}/api/v1/project/list`, { params: params }) }

// delete project
export const deleteProject = params => { return axios.get(`${BaseURL}/api/v1/project/delete`, { params: params }) }

// add project
export const addProject = params => { return axios.post(`${BaseURL}/api/v1/project/add`, qs.stringify(params)) }

// update project
export const updateProject = params => { return axios.post(`${BaseURL}/api/v1/project/update`, qs.stringify(params)) }
