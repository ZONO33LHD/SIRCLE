import React from "react";

export default async function page() {
  const { data } = await client.query({ query: GET_ARTICLES });
  return (
    <div>
      <h1>記事一覧</h1>
      {data.articles.map((article) => (
        <div key={article.id}>
          <h2>{article.title}</h2>
          <p>{article.content}</p>
        </div>
      ))}
    </div>
  );
}
