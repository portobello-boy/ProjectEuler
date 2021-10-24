import setuptools

setuptools.setup(
    name='MathLib',
    version='1.0.0',
    author='Daniel Millson',
    author_email='danielmillson923@gmail.com',
    description='A small package with common tools for solving Project Euler problems',
    long_description='A small package with common tools for solving Project Euler problems',
    long_description_content_type='text/markdown',
    packages=setuptools.find_packages(),
    classifiers=[
        'Programming Language :: Python :: 3',
        'License :: OSI Approved :: MIT License',
        'Operating System :: OS Independent',
    ],
    python_requires='>=3.6',
)