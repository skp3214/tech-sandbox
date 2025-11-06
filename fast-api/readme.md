# FastAPI Crash Course Learning

This folder contains my learning journey with FastAPI, a modern, fast web framework for building APIs with Python.

## 📚 What I'm Learning

FastAPI is a high-performance web framework for building APIs with Python based on standard Python type hints. Key features include:

- **Fast**: Very high performance, on par with NodeJS and Go
- **Fast to code**: Increase the speed to develop features by about 200% to 300%
- **Fewer bugs**: Reduce about 40% of human (developer) induced errors
- **Intuitive**: Great editor support with auto-completion
- **Easy**: Designed to be easy to use and learn
- **Short**: Minimize code duplication
- **Robust**: Get production-ready code with automatic interactive documentation

## 🚀 Project Structure

```
fast-api/
├── main.py              # Main FastAPI application
├── requirements.txt     # Python dependencies
└── readme.md           # This file
```

## 🛠️ What I've Built

### Tea Collection API

A simple CRUD (Create, Read, Update, Delete) API for managing a tea collection. This project demonstrates:

#### **Data Model**
- Using Pydantic for data validation
- Type hints for better code quality
- BaseModel for structured data

#### **API Endpoints**

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Welcome message |
| GET | `/tea` | Get all teas |
| POST | `/tea` | Add a new tea |
| PUT | `/tea/{tea_id}` | Update a specific tea |
| DELETE | `/tea/{tea_id}` | Delete a specific tea |

#### **Features Implemented**
- ✅ **GET**: Retrieve all teas
- ✅ **POST**: Add new tea with validation
- ✅ **PUT**: Update existing tea by ID
- ✅ **DELETE**: Remove tea by ID
- ✅ **Pydantic Models**: Data validation and serialization
- ✅ **Type Hints**: Better code documentation and IDE support

## 🔧 Setup & Installation

### Prerequisites
- Python 3.7+
- pip (Python package installer)

### Installation Steps

1. **Clone the repository** (if not already done):
   ```bash
   git clone https://github.com/skp3214/tech-sandbox.git
   cd tech-sandbox/fast-api
   ```

2. **Create a virtual environment** (recommended):
   ```bash
   python -m venv venv
   venv\Scripts\activate  # On Windows
   ```

3. **Install dependencies**:
   ```bash
   pip install -r requirements.txt
   ```

4. **Run the application**:
   ```bash
   uvicorn main:app --reload
   ```

5. **Access the API**:
   - API: http://127.0.0.1:8000
   - Interactive docs: http://127.0.0.1:8000/docs
   - Alternative docs: http://127.0.0.1:8000/redoc

## 📖 Key Dependencies

- **FastAPI**: The web framework
- **Pydantic**: Data validation using Python type hints
- **Uvicorn**: ASGI server for running the application
- **Starlette**: Lightweight ASGI framework (FastAPI is built on top of it)

## 🧪 Testing the API

### Using the Interactive Documentation
FastAPI automatically generates interactive API documentation at `/docs`. You can:
- View all endpoints
- Test endpoints directly in the browser
- See request/response schemas

### Using curl (Command Line)

```bash
# Get all teas
curl -X GET "http://127.0.0.1:8000/tea"

# Add a new tea
curl -X POST "http://127.0.0.1:8000/tea" \
     -H "Content-Type: application/json" \
     -d '{"id": 1, "name": "Earl Grey", "origin": "England"}'

# Update a tea
curl -X PUT "http://127.0.0.1:8000/tea/1" \
     -H "Content-Type: application/json" \
     -d '{"id": 1, "name": "Earl Grey Premium", "origin": "England"}'

# Delete a tea
curl -X DELETE "http://127.0.0.1:8000/tea/1"
```

## 🎯 Learning Objectives Achieved

- [x] Understanding FastAPI basics
- [x] Creating RESTful API endpoints
- [x] Using Pydantic for data validation
- [x] Implementing CRUD operations
- [x] Working with request/response models
- [x] Understanding automatic documentation generation
- [x] Using type hints effectively

## 🚧 Known Issues & Future Improvements

### Current Issues
- **PUT endpoint bug**: Line 30 should return `updated_tea` instead of `update_tea`
- **Data persistence**: Currently using in-memory storage (data is lost on restart)

### Potential Improvements
- [ ] Add database integration (SQLite, PostgreSQL)
- [ ] Implement proper error handling
- [ ] Add authentication and authorization
- [ ] Add input validation and sanitization
- [ ] Implement pagination for large datasets
- [ ] Add unit tests
- [ ] Add logging
- [ ] Dockerize the application

## 📝 Notes

This is a learning project focused on understanding FastAPI fundamentals. The code demonstrates basic concepts and is not production-ready. Key learning points:

1. **Automatic Documentation**: FastAPI generates OpenAPI schema automatically
2. **Type Safety**: Pydantic models ensure data validation
3. **Performance**: ASGI-based for high performance
4. **Developer Experience**: Excellent IDE support and error messages

## 🔗 Useful Resources

- [FastAPI Official Documentation](https://fastapi.tiangolo.com/)
- [Pydantic Documentation](https://pydantic-docs.helpmanual.io/)
- [Python Type Hints](https://docs.python.org/3/library/typing.html)
- [OpenAPI Specification](https://swagger.io/specification/)

---

*This README is part of my tech learning journey. Check out other technologies I'm exploring in the parent directory!*