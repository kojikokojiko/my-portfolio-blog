import React, { useState, useEffect } from 'react';
import axios from '../axiosConfig';
import { useParams } from 'react-router-dom';

interface Article {
  id: string;
  title: string;
  content: string;
}

const ArticleDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [article, setArticle] = useState<Article | null>(null);

  useEffect(() => {
    const fetchArticle = async () => {
      const response = await axios.get<Article>(`/api/blogs/${id}`);
      setArticle(response.data);
    };

    fetchArticle();
  }, [id]);

  if (!article) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <h1>{article.title}</h1>
      <div dangerouslySetInnerHTML={{__html:`${article.content}`}}></div>
    </div>
  );
};

export default ArticleDetail;
