import type { RequestHandler } from './$types';

export const GET: RequestHandler = async ({ url, params }) => {
	const response = await fetch(`http://localhost:8080/api/${params.type}?${url.searchParams.toString()}`).then(response => response.json());
	return new Response(response);
};