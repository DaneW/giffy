import { baseUrl } from '../config/api';

// Assumtion: user is already logged in
const authToken = Math.floor(Math.random() * 10000 + 1);

const headers = {
  "Content-Type": "application/json",
  "X-Giffy-Token": authToken
};

const getGIFs = async (query, offset = 0) => {
  const response = await fetch(`${baseUrl}/gifs?q=${query}&offset=${offset}`, { headers });
  const json = await response.json();
  return { ...json, query };
};

const upVote = async (id) => {
  const response = await fetch(`${baseUrl}/gifs/${id}/upvote`, { headers, method: "PUT" });
  return response.json();
};
  

const downVote = async (id) => {
  const response = await fetch(`${baseUrl}/gifs/${id}/downvote`, { headers, method: "PUT" });
  return response.json();
}

export {
  getGIFs,
  upVote,
  downVote,
};