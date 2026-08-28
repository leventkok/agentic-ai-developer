export interface Todo {
  id: string;
  title: string;
  done: boolean;
}

export interface TodoFormState {
  title: string;
}

export interface WeatherData {
  city: string;
  tempC: number;
  description: string;
}
