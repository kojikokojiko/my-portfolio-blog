import axios from 'axios';

const instance = axios.create({
  baseURL: 'http://tk2-412-47176.vs.sakura.ne.jp',
});

export default instance;
