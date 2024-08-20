'use client';

import React, { useState } from 'react';
import { useMutation, gql } from '@apollo/client';

const CREATE_TRANSACTION = gql`
  mutation CreateTransaction($input: CreateTransactionInput!) {
    createTransaction(input: $input) {
      id
      amount
      type
      category {
        id
        name
      }
      date
    }
  }
`;

type TransactionType = 'income' | 'expense';

const categoryIds = {
  income: {
    SALARY: '1',
    SIDE_INCOME: '2',
    INVESTMENT: '3',
    OTHER_INCOME: '4',
  },
  expense: {
    FOOD: '5',
    RENT: '6',
    UTILITIES: '7',
    TRANSPORTATION: '8',
    COMMUNICATION: '9',
    ENTERTAINMENT: '10',
    MEDICAL: '11',
    OTHER_EXPENSE: '12',
  },
};

export default function TransactionForm() {
  const [type, setType] = useState<TransactionType>('expense');
  const [amount, setAmount] = useState('');
  const [category, setCategory] = useState({ id: '', name: '' });
  const [date, setDate] = useState('');
  const [createTransaction] = useMutation(CREATE_TRANSACTION);

  const getUserId = () => {
    if (typeof window !== 'undefined') {
      return localStorage.getItem('userId');
    }
    return null;
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!amount || !category.id || !date) {
      alert('全ての項目を入力してください。');
      return;
    }

    const userId = getUserId();
    if (!userId) {
      alert('ユーザーIDが見つかりません。再度ログインしてください。');
      return;
    }

    const typeKey = type.toUpperCase();

    try {
      const result = await createTransaction({
        variables: {
          input: {
            userId: userId,
            amount: parseFloat(amount),
            type: typeKey,
            categoryId: category.id,
            categoryName: category.name,
            date: new Date(date).toISOString(),
            description: '',
            isRecurring: false,
          },
        },
      });

      console.log('取引が正常に保存されました:', result.data.createTransaction);

      // フォームをリセット
      setAmount('');
      setCategory({ id: '', name: '' });
      setDate('');

      alert('取引が正常に保存されました。');
    } catch (error) {
      console.error('取引の保存中にエラーが発生しました:', error);
      alert('取引の保存中にエラーが発生しました。');
    }
  };

  return (
    <div className="rounded-lg bg-white p-6 shadow-md">
      <h2 className="mb-6 text-xl font-bold">取引入力</h2>
      <form onSubmit={handleSubmit}>
        <div className="mb-4">
          <label className="mb-2 block text-sm font-medium text-gray-700">種類</label>
          <div className="flex space-x-2">
            {['income', 'expense'].map((t) => (
              <button
                key={t}
                type="button"
                onClick={() => setType(t as TransactionType)}
                className={`flex-1 rounded-md px-4 py-2 text-sm font-medium ${
                  type === t
                    ? 'bg-orange-400 text-white'
                    : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
                }`}
              >
                {t === 'income' ? '収入' : '支出'}
              </button>
            ))}
          </div>
        </div>
        <div className="mb-4">
          <label className="mb-2 block text-sm font-medium text-gray-700">金額</label>
          <input
            type="number"
            value={amount}
            onChange={(e) => setAmount(e.target.value)}
            placeholder="0"
            className="w-full rounded-md border border-gray-300 px-3 py-2"
          />
        </div>
        <div className="mb-4 h-24">
          <label className="mb-2 block text-sm font-medium text-gray-700">カテゴリ</label>
          <div className="grid grid-cols-2 gap-2 sm:grid-cols-3 md:grid-cols-4">
            {Object.entries(categoryIds[type]).map(([id, name]) => (
              <button
                key={id}
                type="button"
                onClick={() => setCategory({ id: id, name: name })}
                className={`rounded-md px-3 py-2 text-sm font-medium ${
                  category.id === id
                    ? 'bg-green-500 text-white'
                    : 'bg-gray-200 text-gray-700 hover:bg-gray-300'
                }`}
              >
                {name}
              </button>
            ))}
          </div>
        </div>
        <div className="mb-4">
          <label className="mb-2 block text-sm font-medium text-gray-700">日付</label>
          <input
            type="date"
            value={date}
            onChange={(e) => setDate(e.target.value)}
            className="w-full rounded-md border border-gray-300 px-3 py-2"
          />
        </div>
        <button
          type="submit"
          className="w-full rounded-md bg-blue-500 px-4 py-2 text-white hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        >
          保存
        </button>
      </form>
    </div>
  );
}
