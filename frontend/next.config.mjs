const isGitHubPages = process.env.GITHUB_PAGES === 'true';
const repositoryName = process.env.GITHUB_REPOSITORY?.split('/')[1] ?? 'Family-Story-Universe';

/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  ...(isGitHubPages
    ? {
        output: 'export',
        images: {
          unoptimized: true,
        },
        basePath: `/${repositoryName}`,
        assetPrefix: `/${repositoryName}/`,
      }
    : {}),
};

export default nextConfig;
