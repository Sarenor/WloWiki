import type { RequestHandler } from './$types';

export const GET: RequestHandler = async ({ url, params }) => {
	return await fetch(`http://localhost:8080/${params.type}?${url.searchParams.toString()}`);
};